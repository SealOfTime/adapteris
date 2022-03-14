package main

import (
	"fmt"

	"github.com/sealoftime/adapteris/app"
)

func main() {
	cfg := app.NewConfig()
	cfg.Setup()
	fmt.Printf("Config: %+v\n", cfg)

	a := app.New(cfg)
	if err := a.Start(); err != nil {
		fmt.Println(err)
	}
}
