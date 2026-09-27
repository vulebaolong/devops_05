package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	DATABASE_URL := os.Getenv("DATABASE_URL")
	log.Println("DATABASE_URL=", DATABASE_URL)
	if DATABASE_URL == "" {
		log.Println("WARNING: DATABASE_URL is not configured")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(jsonRecovery())
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go Gin Web Framework!",
		})
	})
	router.GET("/user", func(c *gin.Context) {
		db, err := sql.Open("postgres", DATABASE_URL)
		if err != nil {
			writeError(c, http.StatusServiceUnavailable, "DB_CONNECTION_ERROR", "Cannot initialize the database client.", err)
			return
		}
		defer db.Close()
		rows, err := db.Query("SELECT name FROM users")

		if err != nil {
			writeDatabaseError(c, err)
			return
		}
		defer rows.Close()

		users := make([]string, 0)

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				writeError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Cannot read user data.", err)
				return
			}
			users = append(users, name)
		}
		if err := rows.Err(); err != nil {
			writeDatabaseError(c, err)
			return
		}
		if len(users) == 0 {
			writeError(c, http.StatusNotFound, "NO_DATA", "No users found.", nil)
			return
		}

		c.JSON(200, gin.H{
			"source_code": "go",
			"users":       users,
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}

func writeDatabaseError(c *gin.Context, err error) {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "42P01" {
		writeError(c, http.StatusServiceUnavailable, "TABLE_NOT_FOUND", "The database table 'users' does not exist.", err)
		return
	}
	message := strings.ToLower(err.Error())
	if (errors.As(err, &pqErr) && (strings.HasPrefix(string(pqErr.Code), "08") || pqErr.Code == "28P01" || pqErr.Code == "3D000")) || strings.Contains(message, "connect") || strings.Contains(message, "timeout") || strings.Contains(message, "connection refused") {
		writeError(c, http.StatusServiceUnavailable, "DB_CONNECTION_ERROR", "Cannot connect to the database.", err)
		return
	}
	writeError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An unexpected database error occurred.", err)
}

func writeError(c *gin.Context, status int, code, message string, err error) {
	details := ""
	if err != nil {
		details = err.Error()
		log.Printf("request failed: method=%s path=%s code=%s error=%+v", c.Request.Method, c.Request.URL.Path, code, err)
	}
	c.JSON(status, gin.H{"source_code": "go", "success": false, "error": gin.H{"code": code, "message": message, "details": details}, "timestamp": time.Now().UTC(), "path": c.Request.URL.Path})
}

func jsonRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("panic while processing %s %s: %+v", c.Request.Method, c.Request.URL.Path, recovered)
		writeError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An unexpected error occurred.", nil)
	})
}
