package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gilwong00/go-discord/api"
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/pkg/config"
	"github.com/gilwong00/go-discord/pkg/token"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
		os.Exit(1)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
		os.Exit(1)
	}
	store := db.NewStore(conn)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("failed to create token maker", err)
		os.Exit(1)
	}
	server := api.NewApiServer(config, store, tokenMaker)
	server.StartAPIServer()
}
