package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/becardine/learning-go/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDB struct {
	conn *sql.DB
	*db.Queries
}

func NewCourseDB(conn *sql.DB) *CourseDB {
	return &CourseDB{
		conn:    conn,
		Queries: db.New(conn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

// função para executar transações no banco de dados
func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	// nil é o nível de isolamento padrão
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		// se ocorrer um erro, o rollback é feito
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", rbErr, err)
		}
		return err
	}

	// se tudo ocorrer bem, o commit é feito
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, course CourseParams, category CategoryParams) error {
	return c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  category.ID,
			Price:       100.05,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go programming language", Valid: true},
	}

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	}

	courseDB := NewCourseDB(conn)

	err = courseDB.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	if err != nil {
		panic(err)
	}

}
