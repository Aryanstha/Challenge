package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	dsn := "postgresql://aryan:wpfAiBzl6DUz3b97cSWNAw@aryanstha-6390.6xw.aws-ap-southeast-1.cockroachlabs.cloud:26257/license?sslmode=verify-full"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
