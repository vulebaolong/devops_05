using DotNetEnv;
using Npgsql;

Env.Load();

Console.WriteLine($"DB_HOST={Environment.GetEnvironmentVariable("DB_HOST")}");
Console.WriteLine($"DB_PORT={Environment.GetEnvironmentVariable("DB_PORT")}");
Console.WriteLine($"DB_NAME={Environment.GetEnvironmentVariable("DB_NAME")}");
Console.WriteLine($"DB_USER={Environment.GetEnvironmentVariable("DB_USER")}");
Console.WriteLine($"DB_PASSWORD={Environment.GetEnvironmentVariable("DB_PASSWORD")}");
Console.WriteLine($"DB_SSL_MODE={Environment.GetEnvironmentVariable("DB_SSL_MODE")}");
Console.WriteLine($"DATABASE_URL=Host={Environment.GetEnvironmentVariable("DB_HOST")};Port={Environment.GetEnvironmentVariable("DB_PORT") ?? "5432"};Database={Environment.GetEnvironmentVariable("DB_NAME")};Username={Environment.GetEnvironmentVariable("DB_USER")};Password={Environment.GetEnvironmentVariable("DB_PASSWORD")};SSL Mode={Environment.GetEnvironmentVariable("DB_SSL_MODE") ?? "Disable"}");

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddCors(options =>
{
    options.AddDefaultPolicy(policy =>
        policy.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader());
});

var app = builder.Build();
app.UseCors();

app.Use(async (context, next) =>
{
    var logger = context.RequestServices.GetRequiredService<ILogger<Program>>();
    var startedAt = DateTime.UtcNow;
    logger.LogInformation("HTTP request started: {Method} {Path}", context.Request.Method, context.Request.Path);

    try
    {
        await next();
    }
    catch (Exception exception)
    {
        logger.LogError(exception, "Unhandled error while processing {Method} {Path}", context.Request.Method, context.Request.Path);
        if (!context.Response.HasStarted)
        {
            context.Response.StatusCode = StatusCodes.Status500InternalServerError;
            await context.Response.WriteAsJsonAsync(ErrorResponse(
                "INTERNAL_SERVER_ERROR", "An unexpected error occurred.", exception.Message, context.Request.Path));
        }
    }
    finally
    {
        logger.LogInformation("HTTP request completed: {Method} {Path} - {StatusCode} in {ElapsedMs}ms",
            context.Request.Method, context.Request.Path, context.Response.StatusCode,
            (DateTime.UtcNow - startedAt).TotalMilliseconds);
    }
});

app.MapGet("/", () => Results.Json(new { message = "Hello from ASP.NET Core!" }));

app.MapGet("/user", async (HttpContext context, ILogger<Program> logger) =>
{
    try
    {
        var databaseUrl = new NpgsqlConnectionStringBuilder
        {
            Host = Environment.GetEnvironmentVariable("DB_HOST"),
            Port = int.Parse(Environment.GetEnvironmentVariable("DB_PORT") ?? "5432"),
            Database = Environment.GetEnvironmentVariable("DB_NAME"),
            Username = Environment.GetEnvironmentVariable("DB_USER"),
            Password = Environment.GetEnvironmentVariable("DB_PASSWORD"),
            SslMode = Enum.Parse<SslMode>(Environment.GetEnvironmentVariable("DB_SSL_MODE") ?? "Disable")
        }.ConnectionString;
        var users = new List<string>();
        await using var connection = new NpgsqlConnection(databaseUrl);
        await connection.OpenAsync(context.RequestAborted);
        await using var command = new NpgsqlCommand("SELECT name FROM users", connection);
        await using var reader = await command.ExecuteReaderAsync(context.RequestAborted);

        while (await reader.ReadAsync(context.RequestAborted))
            users.Add(reader.GetString(0));

        if (users.Count == 0)
            return Results.Json(ErrorResponse("NO_DATA", "No users found.", null, context.Request.Path), statusCode: 404);

        return Results.Json(new { source_code = "c#", users });
    }
    catch (PostgresException exception) when (exception.SqlState == PostgresErrorCodes.UndefinedTable)
    {
        logger.LogError(exception, "PostgreSQL table 'users' does not exist");
        return Results.Json(ErrorResponse("TABLE_NOT_FOUND", "The database table 'users' does not exist.", exception.MessageText, context.Request.Path), statusCode: 503);
    }
    catch (Exception exception) when ((exception is NpgsqlException && exception is not PostgresException) || exception is TimeoutException)
    {
        logger.LogError(exception, "Cannot connect to or query PostgreSQL");
        return Results.Json(ErrorResponse("DB_CONNECTION_ERROR", "Cannot connect to the database.", exception.Message, context.Request.Path), statusCode: 503);
    }
    catch (Exception exception) when (exception is ArgumentException or FormatException)
    {
        logger.LogError(exception, "Invalid PostgreSQL environment configuration");
        return Results.Json(ErrorResponse("DB_CONFIGURATION_ERROR", "The database configuration is invalid.", exception.Message, context.Request.Path), statusCode: 503);
    }
});

app.Run();

static object ErrorResponse(string code, string message, string? details, PathString path) => new
{
    source_code = "c#",
    success = false,
    error = new { code, message, details },
    timestamp = DateTime.UtcNow,
    path = path.Value
};
