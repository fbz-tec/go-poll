package main

import (
	"context"
	"log"

	"github.com/fbz-tec/go-poll/api"
	db "github.com/fbz-tec/go-poll/db/sqlc"
	"github.com/fbz-tec/go-poll/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}

	log.Printf("Starting the application %s , env: %s\n", "app", config.Environment)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %s", err)
	}

	store := db.New(connPool)

	runGinServer(config, store)
}

func runGinServer(config util.Config, store *db.Queries) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
