package pgx

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func mustConnect(t testing.TB, connString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to establish connection: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(context.Background()); err != nil {
			t.Fatalf("couldn't close connection: %+v\n", err)
		}
	})

	return conn
}
