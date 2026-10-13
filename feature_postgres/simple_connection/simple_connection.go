package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// "postgres://postgres:12345@localhost:5432/postgres"
func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
