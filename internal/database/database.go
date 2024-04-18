package database

import (
	"database/sql"
	"fmt"
	"github.com/leonardogregoriocs/task_management_system/internal/config"
	_ "github.com/lib/pq"
	"log"
)

func Connect(postgresCfg config.PostgresConfig) (*sql.DB, error) {
	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresCfg.Host, postgresCfg.Port, postgresCfg.User, postgresCfg.Password, postgresCfg.Database, "disable",
	)

	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("create connection with postgres with success!")

	return conn, nil
}
