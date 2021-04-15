package main

import (
	"flag"
	"opb/config"
	"opb/lib"
	"opb/routes"
)

func main() {
	env := flag.String("env", "development", "environments flag")
	config.Init(*env)

	db := lib.GetDB()
	defer db.Close()
	lib.InitDB(db)

	router := routes.Init(db)
	router.Run(":3000")
}
