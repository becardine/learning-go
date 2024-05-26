package main

import (
	"fmt"

	"github.com/becardine/utils-secret/pkg/events"
)

func main() {

	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
