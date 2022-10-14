package main

import (
	"log"
	"twittor/db"
	"twittor/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sem conexão com banco de dados")
		return
	}
	// db.ConectarDB()
	handlers.Handlers()
}
