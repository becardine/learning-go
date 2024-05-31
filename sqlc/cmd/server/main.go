package main

import (
	"context"
	"database/sql"

	"github.com/becardine/learning-go/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go programming language", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "5f421ec2-39fb-4df0-8e55-dffcca081739",
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	})

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
