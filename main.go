package main

import (
	"context"
	//"database/sql"
	"log"

	"github.com/naviscom/catalystx/api"
	db "github.com/naviscom/catalystx/db/sqlc"
	"github.com/naviscom/catalystx/util"
	//_ "github.com/lib/pq"
	"github.com/jackc/pgx/v5/pgxpool"
	
)

func main() {
	// Loading environment variables from app.env into config variable of type util.Config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Creatinhg new instance of db connection handle
//	conn, err := sql.Open(config.DBDriver, config.DBSource)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// Associating db connection handle with the Queries interface
//	store := db.NewStore(conn)
	store := db.NewStore(connPool)
	// Creating new instance of HTTP server
	server, _ := api.NewServer(config, store)
	// Giving a kick start to the newly created HTTP server instance
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
