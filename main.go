package main

import (
	"padiplace_ijs/database"
	"padiplace_ijs/router"
)

func main() {
	database.InitialMigration()
	router.InitializeRouter()
}
