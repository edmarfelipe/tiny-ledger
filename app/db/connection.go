package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func getEnv(key string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if len(value) == 0 {
		panic(fmt.Errorf("%s is empty", key))
	}
	return value
}

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", getEnv("DB_USER"), getEnv("DB_PASS"), getEnv("DB_HOST"), getEnv("DB_NAME"))
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
