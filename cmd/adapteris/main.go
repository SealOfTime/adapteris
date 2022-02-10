package main

import (
	"fmt"

	"github.com/sealoftime/adapteris/setup"
)

func main() {
	app := setup.App()
	if err := app.Start(); err != nil {
		fmt.Println(err)
	}
}
