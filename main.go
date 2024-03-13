package main

import (
	"log"

	dbconfig "github.com/KAMM-FARM-SUPPLY/KAMM-Reporter/db_config"
	"github.com/KAMM-FARM-SUPPLY/KAMM-Reporter/internal"
)

func main() {

	//Loading the environment variables
	env := internal.EnvConfig{}
	env_err := env.LoadConfig()
	if env_err != nil {
		log.Fatalf(env_err.Error())
	}

	//Creating a new mongo database instance
	db, db_err := dbconfig.NewDatabaseClient()
	if db_err != nil {
		log.Fatalf(db_err.Error())
	}

	//Connecting to the mongo database instance
	db_client, db_ctx, db_cancet_ctx, db_err2 := db.Connect(env.MONGOURI)
	if db_err2 != nil {
		log.Fatalf(db_err2.Error())
	}

	//Pinging the mongo database client to verify the connection
	db_ping_err := db.Ping(db_client, db_ctx)
	if db_ping_err != nil {
		log.Fatalf(db_ping_err.Error())
	}

	//Closing the connection when the application returns
	defer db.Close(db_client, db_ctx, db_cancet_ctx)

}
