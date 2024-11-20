package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// User represents a registered user
type User struct {
	ID       string
	Password string // Hashed password
}

// Product represents a product with a name, category, price, and quantity
type Product struct {
	Name     string
	Category string
	Price    float64
	Quantity int
}

// UserStore holds all registered users
type UserStore struct {
	mu    sync.Mutex
	users map[string]User
}

// ProductStore holds all products
type ProductStore struct {
	mu       sync.Mutex
	products []Product
}

// NewUserStore initializes a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]User),
	}
}

// NewProductStore initializes a new product store
func NewProductStore() *ProductStore {
	return &ProductStore{
		products: []Product{},
	}
}

// HashPassword securely hashes a plain-text password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// RegisterUser registers a new user with an ID and password
func (store *UserStore) RegisterUser(id, password string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	// Check if the user ID already exists
	if _, exists := store.users[id]; exists {
		return errors.New("user ID already exists")
	}

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Store the user
	store.users[id] = User{
		ID:       id,
		Password: hashedPassword,
	}
	return nil
}

// AddProduct adds a new product to the store
func (store *ProductStore) AddProduct(name, category string, price float64, quantity int) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.products = append(store.products, Product{
		Name:     name,
		Category: category,
		Price:    price,
		Quantity: quantity,
	})
}

// ListProducts displays all products
func (store *ProductStore) ListProducts() {
	store.mu.Lock()
	defer store.mu.Unlock()

	fmt.Println("\n--- Products ---")
	for i, product := range store.products {
		fmt.Printf("%d. Name: %s, Category: %s, Price: %.2f, Quantity: %d\n", i+1, product.Name, product.Category, product.Price, product.Quantity)
	}
}

// EditProductPrice updates the price of a specific product
func (store *ProductStore) EditProductPrice(index int, newPrice float64) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if index < 0 || index >= len(store.products) {
		return errors.New("invalid product index")
	}

	store.products[index].Price = newPrice
	return nil
}

// EditProductQuantity updates the quantity of a specific product
func (store *ProductStore) EditProductQuantity(index int, newQuantity int) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if index < 0 || index >= len(store.products) {
		return errors.New("invalid product index")
	}

	store.products[index].Quantity = newQuantity
	return nil
}

// PromptUserInput gets input from the user
func PromptUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	userStore := NewUserStore()
	productStore := NewProductStore()

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Register a New User")
		fmt.Println("2. Add a New Product")
		fmt.Println("3. List All Products")
		fmt.Println("4. Edit Product Details")
		fmt.Println("5. Exit")
		choice := PromptUserInput("Enter your choice: ")

		switch choice {
		case "1":
			fmt.Println("\n--- Register a New User ---")
			id := PromptUserInput("Enter a new ID: ")
			if id == "" {
				fmt.Println("ID cannot be empty. Please try again.")
				continue
			}

			password := PromptUserInput("Enter a password: ")
			if password == "" {
				fmt.Println("Password cannot be empty. Please try again.")
				continue
			}

			err := userStore.RegisterUser(id, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User successfully registered:", id)
			}

		case "2":
			fmt.Println("\n--- Add a New Product ---")
			name := PromptUserInput("Enter product name: ")
			if name == "" {
				fmt.Println("Product name cannot be empty. Please try again.")
				continue
			}

			category := PromptUserInput("Enter product category: ")
			if category == "" {
				fmt.Println("Category cannot be empty. Please try again.")
				continue
			}

			priceInput := PromptUserInput("Enter product price: ")
			price, err := strconv.ParseFloat(priceInput, 64)
			if err != nil || price < 0 {
				fmt.Println("Invalid price. Please try again.")
				continue
			}

			quantityInput := PromptUserInput("Enter product quantity: ")
			quantity, err := strconv.Atoi(quantityInput)
			if err != nil || quantity < 0 {
				fmt.Println("Invalid quantity. Please try again.")
				continue
			}

			productStore.AddProduct(name, category, price, quantity)
			fmt.Println("Product successfully added:", name)

		case "3":
			productStore.ListProducts()

		case "4":
			productStore.ListProducts()
			indexInput := PromptUserInput("Enter the product number to edit: ")
			index, err := strconv.Atoi(indexInput)
			if err != nil || index < 1 {
				fmt.Println("Invalid product number. Please try again.")
				continue
			}

			fmt.Println("\nSelect what to edit:")
			fmt.Println("1. Edit Price")
			fmt.Println("2. Edit Quantity")
			editChoice := PromptUserInput("Enter your choice: ")

			switch editChoice {
			case "1":
				priceInput := PromptUserInput("Enter new price: ")
				newPrice, err := strconv.ParseFloat(priceInput, 64)
				if err != nil || newPrice < 0 {
					fmt.Println("Invalid price. Please try again.")
					continue
				}

				err = productStore.EditProductPrice(index-1, newPrice)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Product price successfully updated.")
				}

			case "2":
				quantityInput := PromptUserInput("Enter new quantity: ")
				newQuantity, err := strconv.Atoi(quantityInput)
				if err != nil || newQuantity < 0 {
					fmt.Println("Invalid quantity. Please try again.")
					continue
				}

				err = productStore.EditProductQuantity(index-1, newQuantity)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Product quantity successfully updated.")
				}

			default:
				fmt.Println("Invalid choice. Please try again.")
			}

		case "5":
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
