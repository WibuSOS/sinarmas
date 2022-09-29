package main

import (
	"github.com/WibuSOS/sinarmas/api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
