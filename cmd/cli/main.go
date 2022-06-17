package main

import (
	"fmt"

	"github.com/fhott/gocli/commands/start"
	"github.com/fhott/gocli/internal"
)

func main() {
	list := []internal.Command{
		&start.Start{},
	}

	if err := internal.CommandInit("gocli").Start(list); err != nil {
		fmt.Println(err)
	}
}
