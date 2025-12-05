package main

import (
	"log"

	"github.com/vanhieuhp/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := application{
		config: cfg,
	}

	log.Println("Starting server...")

	mux := app.mount()
	err := app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
