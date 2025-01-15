package main

import (
	"fmt"

	"github.com/IlliaSh1/backend/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Println("new app error:", err)
		return
	}

	err = app.Run()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
