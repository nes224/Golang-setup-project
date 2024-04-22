package main

import (
	"context"
	"log"

	"github.com/abs/bestinter/internal/api"
	db "github.com/abs/bestinter/db/sqlc"
	"github.com/abs/bestinter/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot")
	}
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	store := db.NewStore(connPool)
	runGinServer(config, store)
}
