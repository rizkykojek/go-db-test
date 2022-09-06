package golang_db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id,name) VALUES('2', 'kojek')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)

		fmt.Println("ID = ", id)
		fmt.Println("Name = ", name)
	}
}

func TestQuerySqlComplete(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var balance int32
		var rating float32
		var married bool
		var createdAt, birthDate time.Time
		rows.Scan(&id, &name, &balance, &rating, &birthDate, &married, &createdAt)
		fmt.Println("ID = ", id, ", Name = ", name, ", Balance = ", balance, ", birth date = ", birthDate, ", married = ", married, ", created_at = ", createdAt)
	}
}
