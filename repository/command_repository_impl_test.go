package repository

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	golang_db "golang-db"
	"golang-db/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	db := golang_db.GetConnection()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(tx)
	comment := entity.Comment{
		Email:   "kojek2@gmail.com",
		Comment: "OK",
	}

	ctx := context.Background()
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	db := golang_db.GetConnection()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(tx)

	ctx := context.Background()
	id := int32(6)
	result, err := commentRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	db := golang_db.GetConnection()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(tx)

	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range result {
		fmt.Println(comment)
	}
	tx.Commit()
}
