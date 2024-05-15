package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// cria um contexto com cancelamento automático após 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	// caso o contexto seja cancelado, o case abaixo é executado
	case <-ctx.Done():
		fmt.Println("hotel booking cancelled. timeout reached")
		return
	// caso o contexto ultrapasse o tempo de 2 segundos, o case abaixo é executado
	case <-time.After(5 * time.Second):
		fmt.Println("hotel booking done")
	}
}
