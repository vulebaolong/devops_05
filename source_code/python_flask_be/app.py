import os
from datetime import datetime, timezone

import psycopg
from dotenv import load_dotenv
from flask import Flask, jsonify, request

load_dotenv()
print("DATABASE_URL=", os.getenv("DATABASE_URL", ""), flush=True)
app = Flask(__name__)


@app.before_request
def handle_options():
    if request.method == "OPTIONS":
        return "", 204


@app.after_request
def add_cors_headers(response):
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
    response.headers["Access-Control-Allow-Headers"] = "*"
    return response


def write_error(status, code, message):
    return jsonify({
        "source_code": "python",
        "success": False,
        "error": {
            "code": code,
            "message": message,
            "details": ""
        },
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "path": request.path
    }), status


@app.get("/")
def hello():
    return jsonify({"message": "Hello from Python Flask Web Framework!"})


@app.get("/user")
def get_users():
    database_url = os.getenv("DATABASE_URL")
    if not database_url:
        return write_error(503, "DB_CONNECTION_ERROR", "DATABASE_URL is not configured.")

    try:
        with psycopg.connect(database_url, connect_timeout=5) as connection:
            with connection.cursor() as cursor:
                cursor.execute("SELECT name FROM users")
                rows = cursor.fetchall()

        users = []
        for row in rows:
            users.append(row[0])

        if len(users) == 0:
            return write_error(404, "NO_DATA", "No users found.")

        return jsonify({"source_code": "python", "users": users})

    except psycopg.errors.UndefinedTable:
        return write_error(503, "TABLE_NOT_FOUND", "The database table 'users' does not exist.")
    except (psycopg.OperationalError, psycopg.ProgrammingError) as error:
        if isinstance(error, psycopg.OperationalError) or error.sqlstate is None:
            return write_error(503, "DB_CONNECTION_ERROR", "Cannot connect to the database.")
        app.logger.error("Database query failed: %s", type(error).__name__)
        return write_error(500, "INTERNAL_SERVER_ERROR", "An unexpected database error occurred.")


@app.errorhandler(500)
def internal_error(error):
    return write_error(500, "INTERNAL_SERVER_ERROR", "An unexpected error occurred.")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
