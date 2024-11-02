package main

import (
	"app/internal/pkg/app"
	"log"
)

func main() {
	test, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := test.Run("8000"); err != nil {
		log.Fatal(err)
	}
}