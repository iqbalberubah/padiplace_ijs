package main

import (
	d "padiplace_ijs/database"
	r "padiplace_ijs/router"
)

func main() {
	d.InitialMigration()
	r.InitializeRouter()
}
