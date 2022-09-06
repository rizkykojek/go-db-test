package golang_db

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("pgx", "postgresql://localhost/go-db?user=rizkymaulana&password=")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestOpenPoolConnection(t *testing.T) {
	db, err := pgxpool.Connect(context.Background(), "postgresql://localhost/go-db?user=rizkymaulana&password=")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
