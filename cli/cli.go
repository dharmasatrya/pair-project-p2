package cli

import (
	"fmt"
	"game/handler"
	"log"
)

type CLI struct {
	Handler handler.Handler
}

func NewCLI(handler handler.Handler) *CLI {
	return &CLI{
		Handler: handler,
	}
}

func (c *CLI) Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// show initial menu
	c.showMenu()
}

func (c *CLI) showMenu() {
	fmt.Println("Select report to generate:")
	fmt.Println("1. function report")
	fmt.Println("2. function crud")
	fmt.Println("3. Exit")

	var choice int
	fmt.Print("Enter the number of the report you want to generate: (1/2/3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		c.sebuahFucntionReport()
	case 2:
		c.sebuahFucntionCRUD()
	case 3:
		fmt.Println("bubye!")
		return
	default:
		fmt.Println("Invalid choice")
	}

	c.showMenu()
}

func (c *CLI) sebuahFucntionReport() {
	err := c.Handler.SebuahFucntionReport()
	if err != nil {
		log.Print("Error listing game sales: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) sebuahFucntionCRUD() {
	err := c.Handler.SebuahFucntionCRUD()
	if err != nil {
		log.Print("Error listing most popular games: ", err)
		log.Fatal(err)
	}
}
