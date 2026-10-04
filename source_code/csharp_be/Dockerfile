FROM mcr.microsoft.com/dotnet/sdk:10.0-alpine AS build

WORKDIR /app

COPY csharp-be.csproj .
RUN dotnet restore

COPY . .
RUN dotnet publish -c Release -o /app/publish --no-restore

FROM mcr.microsoft.com/dotnet/sdk:10.0-alpine

WORKDIR /app

COPY --from=build /app/publish .

CMD [ "dotnet", "csharp-be.dll" ]

# docker build -t csharp_be_img .
# docker run --name csharp_be -d -p 8001:8080 --env-file=.env csharp_be_img
