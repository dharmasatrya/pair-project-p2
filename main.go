// Go & MySQL CRUD Example
package main

import (
	"game/cli"
	"game/config"
	"game/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectDB()

	handler := handler.NewHandler(db)

	cli := cli.NewCLI(handler)
	cli.Init()
}
