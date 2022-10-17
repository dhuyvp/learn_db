package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"%s",
			os.Getenv("POSTGRES_URL"),
		)
	case "mysql":
		// URL for MySQL connection.
		url = fmt.Sprintf(
			"%s",
			os.Getenv("MYSQL_URL"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s",
			os.Getenv("REDIS_URL"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
