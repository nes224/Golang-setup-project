package main

import (
	"context"
	"log"

	"github.com/abs/bestinter/api"
	db "github.com/abs/bestinter/db/sqlc"
	"github.com/abs/bestinter/util"
	"github.com/jackc/pgx/v5/pgxpool"
)


func runGinServer(config util.Config, store db.Store) {
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	store := db.NewStore(connPool)
	runGinServer(config,store)
}