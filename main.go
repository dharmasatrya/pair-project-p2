// Go & MySQL CRUD Example
package main

import (
	"game/cli"
	"game/config"
	"game/handler"
	"game/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectDB()

	repository := repository.NewRepository(db)

	handler := handler.NewHandler(repository)

	cli := cli.NewCLI(handler)
	cli.Init()
}
