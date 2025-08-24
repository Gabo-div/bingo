package database

import (
	"context"
	"os"

	"github.com/Gabo-div/bingo/packages/database/repositories/go"
	"github.com/jackc/pgx/v5"
)

var connectionInstance *pgx.Conn
var queriesInstance *queries.Queries

func GetConnection() *pgx.Conn {
	if connectionInstance != nil {
		return connectionInstance
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")

	if err != nil {
		os.Exit(1)
	}

	connectionInstance = conn

	return conn
}

func GetQueries() *queries.Queries {
	if queriesInstance != nil {
		return queriesInstance
	}

	conn := GetConnection()

	q := queries.New(conn)

	queriesInstance = q

	return q
}
