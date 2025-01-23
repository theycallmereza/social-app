package main

import (
	"github.com/theycallmereza/social-app/internal/db"
	"github.com/theycallmereza/social-app/internal/env"
	"github.com/theycallmereza/social-app/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetEnvSString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetEnvSString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetEnvInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnvSString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	defer db.Close()
	log.Println("Database connection established")

	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
