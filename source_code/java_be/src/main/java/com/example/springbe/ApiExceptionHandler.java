package com.example.springbe;

import jakarta.servlet.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.dao.DataAccessException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.sql.SQLException;
import java.time.Instant;
import java.util.LinkedHashMap;
import java.util.Map;

@RestControllerAdvice
public class ApiExceptionHandler {
    private static final Logger log = LoggerFactory.getLogger(ApiExceptionHandler.class);

    @ExceptionHandler(NoDataException.class)
    public ResponseEntity<Map<String, Object>> noData(NoDataException exception, HttpServletRequest request) {
        return response(HttpStatus.NOT_FOUND, "NO_DATA", exception.getMessage(), null, request);
    }

    @ExceptionHandler(DataAccessException.class)
    public ResponseEntity<Map<String, Object>> database(DataAccessException exception, HttpServletRequest request) {
        log.error("Database error while processing {} {}", request.getMethod(), request.getRequestURI(), exception);
        SQLException sqlException = findSqlException(exception);
        if (sqlException != null && "42P01".equals(sqlException.getSQLState())) {
            return response(HttpStatus.SERVICE_UNAVAILABLE, "TABLE_NOT_FOUND",
                    "The database table 'users' does not exist.", sqlException.getMessage(), request);
        }
        String sqlState = sqlException == null ? null : sqlException.getSQLState();
        if (sqlState == null || sqlState.startsWith("08") || "28P01".equals(sqlState) || "3D000".equals(sqlState)) {
            return response(HttpStatus.SERVICE_UNAVAILABLE, "DB_CONNECTION_ERROR",
                    "Cannot connect to the database.", exception.getMostSpecificCause().getMessage(), request);
        }
        return response(HttpStatus.INTERNAL_SERVER_ERROR, "INTERNAL_SERVER_ERROR",
                "An unexpected database error occurred.", exception.getMostSpecificCause().getMessage(), request);
    }

    @ExceptionHandler(SQLException.class)
    public ResponseEntity<Map<String, Object>> sql(SQLException exception, HttpServletRequest request) {
        log.error("Database error while processing {} {}", request.getMethod(), request.getRequestURI(), exception);
        if ("42P01".equals(exception.getSQLState())) {
            return response(HttpStatus.SERVICE_UNAVAILABLE, "TABLE_NOT_FOUND",
                    "The database table 'users' does not exist.", exception.getMessage(), request);
        }
        String sqlState = exception.getSQLState();
        if (sqlState == null || sqlState.startsWith("08") || "28P01".equals(sqlState) || "3D000".equals(sqlState)) {
            return response(HttpStatus.SERVICE_UNAVAILABLE, "DB_CONNECTION_ERROR",
                    "Cannot connect to the database.", exception.getMessage(), request);
        }
        return response(HttpStatus.INTERNAL_SERVER_ERROR, "INTERNAL_SERVER_ERROR",
                "An unexpected database error occurred.", exception.getMessage(), request);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<Map<String, Object>> unknown(Exception exception, HttpServletRequest request) {
        log.error("Unhandled error while processing {} {}", request.getMethod(), request.getRequestURI(), exception);
        return response(HttpStatus.INTERNAL_SERVER_ERROR, "INTERNAL_SERVER_ERROR",
                "An unexpected error occurred.", exception.getMessage(), request);
    }

    private SQLException findSqlException(Throwable throwable) {
        while (throwable != null) {
            if (throwable instanceof SQLException sqlException) return sqlException;
            throwable = throwable.getCause();
        }
        return null;
    }

    private ResponseEntity<Map<String, Object>> response(HttpStatus status, String code, String message,
                                                          String details, HttpServletRequest request) {
        Map<String, Object> error = new LinkedHashMap<>();
        error.put("code", code);
        error.put("message", message);
        error.put("details", details);
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("source_code", "java");
        body.put("success", false);
        body.put("error", error);
        body.put("timestamp", Instant.now());
        body.put("path", request.getRequestURI());
        return ResponseEntity.status(status).body(body);
    }
}
