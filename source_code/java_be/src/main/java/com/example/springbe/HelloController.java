package com.example.springbe;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.ResponseEntity;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@RestController
@CrossOrigin(origins = "*")
public class HelloController {
    @GetMapping("/")
    public Map<String, String> hello() {
        return Map.of("message", "Hello from Spring Boot!");
    }

	@GetMapping("/user")
	public ResponseEntity<Map<String, Object>> users() throws SQLException {
		String databaseUrl = System.getProperty("DATABASE_URL", System.getenv("DATABASE_URL"));
		List<String> users = new ArrayList<>();
		try (Connection connection = DriverManager.getConnection(databaseUrl);
			 Statement statement = connection.createStatement();
			 ResultSet resultSet = statement.executeQuery("SELECT name FROM users")) {
			while (resultSet.next()) {
				users.add(resultSet.getString("name"));
			}
		}
		if (users.isEmpty()) {
			throw new NoDataException("No users found.");
		}
		return ResponseEntity.ok(Map.of("source_code", "java", "users", users));
	}
}
