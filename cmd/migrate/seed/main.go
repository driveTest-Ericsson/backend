package main

import (
	"log"

	"github.com/driveTest-Ericsson/backend/internal/db"
	"github.com/driveTest-Ericsson/backend/internal/env"
	"github.com/driveTest-Ericsson/backend/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)
	db.Seed(store, conn)
}
