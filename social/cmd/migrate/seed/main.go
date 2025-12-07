package main

import (
	"fmt"
	"log"

	"github.com/vanhieuhp/social/internal/db"
	"github.com/vanhieuhp/social/internal/env"
	"github.com/vanhieuhp/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:password@localhost:5432/social?sslmode=disable")
	fmt.Println(addr)
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	storage := store.NewStorage(conn)

	db.Seed(storage)
}
