package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", "gopher")
	bookHotel(ctx)
}

// em qualquer função que receba um context, ele deve ser o primeiro parâmetro (convenção	go)
func bookHotel(ctx context.Context) {
	user := ctx.Value("user")
	if user != nil {
		fmt.Println("found user value in context:", user)
		return
	}
	fmt.Println("user value not found in context")
}
