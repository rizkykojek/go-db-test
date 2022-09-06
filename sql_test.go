package golang_db

import (
	"context"
	"fmt"
	"testing"
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

	script := "SELECT * FROM customer"
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
