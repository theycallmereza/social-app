package main

import (
	"log"

	"github.com/theycallmereza/social-app/internal/db"
	"github.com/theycallmereza/social-app/internal/env"
	"github.com/theycallmereza/social-app/internal/store"
)

func main() {
	addr := env.GetEnvString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
