package main

import (
	"log"

	"github.com/vanhieuhp/social/internal/db"
	"github.com/vanhieuhp/social/internal/env"
	"github.com/vanhieuhp/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:password@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	postgresDb, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer postgresDb.Close()
	log.Println("Database connection pool established.")

	postgresStorage := store.NewStorage(postgresDb)

	app := &application{
		config: cfg,
		store:  postgresStorage,
	}

	log.Println("Starting server...")

	mux := app.mount()
	err = app.run(mux)
	if err != nil {
		log.Fatal(err)
	}
}
