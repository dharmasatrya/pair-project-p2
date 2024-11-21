package cli

import (
	"bufio"
	"fmt"
	"game/handler"
	"log"
	"os"
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
	var choice int

	for {
		fmt.Println("Select function:")
		fmt.Println("1. function crud")
		fmt.Println("2. Reportings")
		fmt.Println("3. Exit")
		fmt.Print("Enter the number of the report you want to generate: (1/2/3): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid option. Please enter a number between 1 and 3.")
			continue
		}
		break
	}

	if choice == 3 {
		fmt.Println("Byebye")
		return
	}

	//CRUD options
	if choice == 1 {
		for {
			fmt.Println("Select function:")
			fmt.Println("1. Update Product Category by ID")
			fmt.Println("2. Buy Product")
			fmt.Println("3. Update Name User by ID")
			fmt.Println("4. Create New Category")
			fmt.Println("5. Exit")
			fmt.Print("Enter the number of the report you want to generate: (1/2/3): ")
			_, err := fmt.Scanln(&choice)
			if err != nil || choice < 1 || choice > 5 {
				fmt.Println("Invalid option. Please enter a number between 1 and 3.")
				fmt.Scanln()
				continue
			}
			break
		}

		switch choice {
		case 1:
			c.updateProductCategoryById()
		case 2:
			c.buyProduct()
		case 3:
			c.updateUserNamerById()
		case 4:
			c.createNewCategory()
		case 5:
			fmt.Println("Thanks For Using this CLI!")
			os.Exit(0)
		}
	}

	//REPORTING options
	if choice == 2 {

		for {
			fmt.Println("Select report to generate:")
			fmt.Println("1. Show Orders")
			fmt.Println("2. Show Users Spending")
			fmt.Println("3. Show Current Stocks")
			fmt.Println("4. Go Back")
			fmt.Println("5. Exit")
			fmt.Print("Enter the number of the report you want to generate: (1/2/3/4/5): ")
			_, err := fmt.Scanln(&choice)
			if err != nil || choice < 1 || choice > 5 {
				fmt.Println("Invalid option. Please enter a number between 1 and 5.")
				continue
			}
			break
		}

		switch choice {
		case 1:
			c.showOrders()
		case 2:
			c.showUserSpending()
		case 3:
			c.showStocks()
		case 4:
			c.showMenu()
		case 5:
			fmt.Println("Thanks For Using this CLI!")
			os.Exit(0)
		}
	}

	c.showMenu()
}

// CRUD FUNCTIONS
func (c *CLI) updateProductCategoryById() {
	var id int
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input the ID of the product that you want to update: ")
	fmt.Scanln(&id)
	fmt.Println("Please input the new Name of the product category")
	if scanner.Scan() {
		name = scanner.Text()
	}
	err := c.Handler.UpdateProductCategoryById(id, name)
	if err != nil {
		log.Print("Error listing most popular games: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) updateUserNamerById() {
	var id int
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	//masukin id user
	fmt.Println("Please input the ID of the user that you want to update")
	fmt.Scanln(&id)

	//masukin nama user yang baru
	fmt.Println("Please input the new Name of user")
	if scanner.Scan() {
		name = scanner.Text()
	}
	err := c.Handler.UpdateUserNameById(id, name)
	if err != nil {
		log.Println("Error update name: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) createNewCategory() {
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input Name Category")
	if scanner.Scan() {
		name = scanner.Text()
		fmt.Println("You entered:", name)
	}
	err := c.Handler.CreateNewCategory(name)
	if err != nil {
		log.Print("Error create new category product: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) buyProduct() {
	var categoryID, productID, userID int

	//masukin id user
	fmt.Println("Please input your ID")
	fmt.Scanln(&userID)

	//munculkan category
	errCategories := c.Handler.ShowProductCategories()
	if errCategories != nil {
		log.Print("Error buying product: ", errCategories)
		log.Fatal(errCategories)
	}

	//munculkan product dari kategori pilihan
	fmt.Println("Please input category id: ")
	fmt.Scanln(&categoryID)
	c.Handler.ShowProductByCategoryId(categoryID)

	//beli product
	fmt.Println("Please input the id of the product")
	fmt.Scanln(&productID)

	err := c.Handler.BuyProduct(productID, userID)
	if err != nil {
		log.Print("Error buying product: ", err)
		log.Fatal(err)
	}
}

// REPORTING FUNCTIONS
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
