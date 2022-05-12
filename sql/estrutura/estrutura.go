package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

	urlDatabase := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), urlDatabase)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	db, err := sql.Open("pg", "postgres:postgres@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn.QueryRow(context.Background(), "create database if not exists cursogo")
	conn.QueryRow(context.Background(), "use cursogo")
	conn.QueryRow(context.Background(), "drop table if exists usuarios")
	conn.QueryRow(context.Background(), `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
