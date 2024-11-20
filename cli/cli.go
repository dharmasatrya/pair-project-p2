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
	fmt.Println("1. show orders")
	fmt.Println("2. function crud")
	fmt.Println("3. Exit")

	var choice int
	fmt.Print("Enter the number of the report you want to generate: (1/2/3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		c.showOrders()
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

func (c *CLI) showUserSpending() {
	err := c.Handler.ShowUserSpending()
	if err != nil {
		log.Print("Error listing user spending: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) showStocks() {
	err := c.Handler.ShowStocks()
	if err != nil {
		log.Print("Error listing product stocks: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) showOrders() {
	var month, year int

	for {
		fmt.Print("Please input the month of the report (1-12): ")
		_, err := fmt.Scanln(&month)
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid month. Please enter a number between 1 and 12.")
			continue
		}
		break
	}

	for {
		fmt.Print("Please input the year of the report (e.g., 2024): ")
		_, err := fmt.Scanln(&year)
		if err != nil || year < 2024 || year > 2024 {
			fmt.Println("Invalid year. Please enter a valid year.")
			continue
		}
		break
	}

	err := c.Handler.ShowOrders(month, year)
	if err != nil {
		log.Printf("Error listing orders for %d/%d: %v", month, year, err)
		return
	}

	fmt.Println("Order report displayed successfully.")
}

func (c *CLI) sebuahFucntionCRUD() {
	err := c.Handler.SebuahFucntionCRUD()
	if err != nil {
		log.Print("Error listing most popular games: ", err)
		log.Fatal(err)
	}
}
