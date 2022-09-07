package golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	paramID := "1"
	paramName := "kojek"
	script := "INSERT INTO customer(id,name) VALUES($1, $2)"
	_, err := db.ExecContext(ctx, script, paramID, paramName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert")
}

func TestInsertSqlUsingPreparedStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id,name) VALUES($1, $2)"
	statement, err := db.PrepareContext(ctx, script)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	for i := 5; i < 10; i++ {
		paramName := "kojek_" + strconv.Itoa(i)
		_, err := statement.ExecContext(ctx, strconv.Itoa(i), paramName)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success insert")
}

func TestInsertSqlUsingTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id,name) VALUES($1, $2)"
	tx, err := db.Begin()

	statement, err := tx.PrepareContext(ctx, script)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	for i := 5; i < 10; i++ {
		paramName := "kojek_" + strconv.Itoa(i)
		_, err := statement.ExecContext(ctx, strconv.Itoa(i), paramName)
		if err != nil {
			panic(err)
		}
	}

	selectTotal := "SELECT count(*) FROM customer"
	var total int64
	row := db.QueryRowContext(ctx, selectTotal)
	row.Scan(&total)
	fmt.Println("Total row before commit = ", total)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	row = db.QueryRowContext(ctx, selectTotal)
	row.Scan(&total)
	fmt.Println("Total row after commit = ", total)

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

func TestQuerySqlWithParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	paramID := "1"
	paramName := "kojek"
	script := "SELECT id, name FROM customer where id = $1 AND name = $2"
	rows, err := db.QueryContext(ctx, script, paramID, paramName)
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

	script := "SELECT id, name, balance, rating, birth_date, married, created_at, email FROM customer"
	rows, err := db.QueryContext(ctx, script)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var married bool
		var createdAt time.Time
		var birthDate sql.NullTime

		rows.Scan(&id, &name, &balance, &rating, &birthDate, &married, &createdAt, &email)
		fmt.Println("================")
		fmt.Println("ID = ", id)
		fmt.Println("Name = ", name)
		if email.Valid {
			fmt.Println("Email = ", email.String)
		}
		fmt.Println("Balance = ", balance)
		if birthDate.Valid {
			fmt.Println("Birth date = ", birthDate.Time)
		}
		fmt.Println("Married = ", married)
		fmt.Println("Created at = ", createdAt)
	}
}
