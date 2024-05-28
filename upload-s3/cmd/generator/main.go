package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	i := 0
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		f.WriteString("Hello, World!")
		i++
	}
}
