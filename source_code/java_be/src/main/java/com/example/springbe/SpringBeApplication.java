package com.example.springbe;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.jdbc.autoconfigure.DataSourceAutoConfiguration;

import io.github.cdimascio.dotenv.Dotenv;

@SpringBootApplication(exclude = DataSourceAutoConfiguration.class)
public class SpringBeApplication {

	public static void main(String[] args) {
		Dotenv dotenv = Dotenv.configure().ignoreIfMissing().load();
		String databaseUrl = dotenv.get("DATABASE_URL", System.getenv("DATABASE_URL"));
		System.out.println("DATABASE_URL=" + databaseUrl);
		if (databaseUrl != null) {
			System.setProperty("DATABASE_URL", databaseUrl);
		}
		SpringApplication.run(SpringBeApplication.class, args);
	}

}
