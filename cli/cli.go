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
	fmt.Println("1. Generate Total Game Sales")
	fmt.Println("2. Show Most Popular Games")
	fmt.Println("3. Show Total Games Revenue")
	fmt.Println("4. Show Total Unique Player")
	fmt.Println("5. Exit")

	var choice int
	fmt.Print("Enter the number of the report you want to generate: (1/2/3/4/5): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		c.showTotalGameSales()
	case 2:
		c.showMostPopularGames()
	case 3:
		c.showTotalGamesRevenue()
	case 4:
		c.showTotalPlayer()
	case 5:
		fmt.Println("bubye!")
		return
	default:
		fmt.Println("Invalid choice")
	}

	c.showMenu()
}

func (c *CLI) showTotalGameSales() {
	err := c.Handler.ShowTotalGameSales()
	if err != nil {
		log.Print("Error listing game sales: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) showMostPopularGames() {
	err := c.Handler.ShowMostPopularGames()
	if err != nil {
		log.Print("Error listing most popular games: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) showTotalGamesRevenue() {
	err := c.Handler.ShowTotalGamesRevenue()
	if err != nil {
		log.Print("Error listing total games revenue: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) showTotalPlayer() {
	err := c.Handler.ShowTotalPlayer()
	if err != nil {
		log.Print("Error listing total games revenue: ", err)
		log.Fatal(err)
	}
}
