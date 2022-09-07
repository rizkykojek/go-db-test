package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-db/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	TX *sql.Tx
}

func NewCommentRepository(tx *sql.Tx) CommentRepository {
	return &commentRepositoryImpl{TX: tx}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comment(email, comment) VALUES($1, $2) RETURNING id"
	var id int32
	err := repository.TX.QueryRowContext(ctx, script, comment.Email, comment.Comment).Scan(&id)
	if err != nil {
		return comment, err
	}

	comment.Id = id
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment where id = $1 LIMIT 1"
	rows, err := repository.TX.QueryContext(ctx, script, id)
	defer rows.Close()
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("ID " + strconv.Itoa(int(id)) + " Not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment"
	rows, err := repository.TX.QueryContext(ctx, script)
	defer rows.Close()
	comments := []entity.Comment{}
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
