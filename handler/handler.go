package handler

import (
	"fmt"
	"game/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Handler is the handler interface
type Handler interface {
	UpdateProductCategoryById(id int, newName string) error
	// UpdateUserNameById(id int, newName string) error
	// CreateNewCategory(name string) error
	// BuyProduct(userID, productID int) error
	// DeleteTransactionById(transactionID int) error
	// ShowUserSpending() error
	// ShowOrders(month, year int) error
	// ShowStocks() error
	// ShowProductByCategoryId(id int) error
	// ShowProductCategories() error
}

// HandlerImpl is the handler implementation
// type HandlerImpl struct {
// 	DB *sql.DB
// }

type HandlerImpl struct {
	Repository repository.GameRepository
}

// NewHandler creates a new handler
// func NewHandler(db *sql.DB) *HandlerImpl {
// 	return &HandlerImpl{
// 		DB: db,
// 	}
// }

func NewHandler(repository repository.GameRepository) *HandlerImpl {
	return &HandlerImpl{
		Repository: repository,
	}
}

// FUNCTIONS BUAT CRUD
func (h *HandlerImpl) UpdateProductCategoryById(id int, newName string) error {

	fmt.Printf("ini id yang ada di handler -> %v", id)

	h.Repository.UpdateProductCategoryById(id, newName)

	fmt.Println("product category name has been changed")

	return nil
// }

// func (h *HandlerImpl) UpdateUserNameById(id int, newName string) error {
// 	rows, err := h.DB.Query(`
// 	UPDATE users
// 	SET name = ?
// 	WHERE userID = ?;
// 	`, newName, id)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("name user has been changed")

// 	return nil
// }

// func (h *HandlerImpl) CreateNewCategory(name string) error {
// 	rows, err := h.DB.Query(`
// 	INSERT INTO product_categories (name)
// 	VALUES (?);`, name)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("You have successfully add new category")

// 	return nil
// }

// func (h *HandlerImpl) BuyProduct(userID, productID int) error {

// 	rows, err := h.DB.Query(`
// 	INSERT INTO transactions (userID, productID, purchasedAt)
// 	VALUES (?, ?, NOW());
// 	`, userID, productID)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("You have successfully purchased the item")

// 	return nil
// }

// func (h *HandlerImpl) DeleteTransactionById(transactionID int) error {

// 	rows, err := h.DB.Query(`
// 		DELETE FROM transactions WHERE trxID = ?;
// 	`, transactionID)
// 	if err != nil {
// 		log.Print("Error deleting records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("You have successfully deleted the transaction")

// 	return nil
// }

// // FUNCTION BUAT NAMPILIN REPORT
// func (h *HandlerImpl) ShowUserSpending() error {
// 	rows, err := h.DB.Query(`
// SELECT 
//     u.name AS user_name,
//     SUM(p.price) AS total_spending
// FROM 
//     transactions t
// JOIN 
//     users u ON t.userID = u.userID
// JOIN 
//     products p ON t.productID = p.productID
// GROUP BY 
//     u.userID, u.name
// ORDER BY 
//     total_spending DESC;
// 	`)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("Total Game Sales Report:")
// 	//looping data
// 	for rows.Next() {
// 		var user_name string
// 		var total_spending int

// 		//check error
// 		err = rows.Scan(&user_name, &total_spending)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}
// 		//print data
// 		fmt.Printf("User Name: %s | Total Money Spent: Rp %d\n",
// 			user_name, total_spending)
// 	}

// 	return nil
// }

// func (h *HandlerImpl) ShowOrders(month, year int) error {

// 	rows, err := h.DB.Query(`
// 		SELECT 
// 			t.trxID AS transaction_id,
// 			u.name AS user_name,
// 			u.email AS user_email,
// 			p.name AS product_name,
// 			p.price AS product_price,
// 			t.purchasedAt AS purchase_date
// 		FROM 
// 			transactions t
// 		JOIN 
// 			users u ON t.userID = u.userID
// 		JOIN 
// 			products p ON t.productID = p.productID
// 		WHERE 
// 			MONTH(t.purchasedAt) = ? AND YEAR(t.purchasedAt) = ?
// 		ORDER BY 
// 			t.purchasedAt ASC;
// 	`, month, year)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("Total Game Sales Report:")

// 	if !rows.Next() {
// 		fmt.Println("No records found on this timerange")
// 	}

// 	for rows.Next() {
// 		var transactionID int
// 		var userName, userEmail, productName string
// 		var productPrice int
// 		var purchaseDate string

// 		err = rows.Scan(&transactionID, &userName, &userEmail, &productName, &productPrice, &purchaseDate)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}

// 		// Print each transaction
// 		fmt.Printf("Transaction ID: %d | User: %s (%s) | Product: %s | Price: Rp %d | Date: %s\n",
// 			transactionID, userName, userEmail, productName, productPrice, purchaseDate)
// 	}

// 	// Check for errors during rows iteration
// 	if err = rows.Err(); err != nil {
// 		log.Print("Error during rows iteration: ", err)
// 		return err
// 	}

// 	return nil
// }

// func (h *HandlerImpl) ShowStocks() error {
// 	rows, err := h.DB.Query(`
// 		SELECT 
// 			pc.name AS category_name,
// 			p.name AS product_name,
// 			p.quantity AS current_quantity
// 		FROM 
// 			products p
// 		JOIN 
// 			product_categories pc ON p.categoryID = pc.categoryID
// 		ORDER BY 
// 			pc.name, p.name;
// 	`)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("Total Game Sales Report:")
// 	//looping data
// 	for rows.Next() {
// 		var category_name, product_name string
// 		var current_quantity int

// 		//check error
// 		err = rows.Scan(&category_name, &product_name, &current_quantity)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}
// 		//print data
// 		fmt.Printf("Category: %s | Product Name: %s | Available Stock: %d\n",
// 			category_name, product_name, current_quantity)
// 	}

// 	return nil
// }

// func (h *HandlerImpl) ShowProductByCategoryId(id int) error {
// 	rows, err := h.DB.Query(`
// 		SELECT 
// 			p.name AS product_name,
// 			p.price AS price,
// 			p.quantity AS stock
// 		FROM 
// 			products p
// 		JOIN 
// 			product_categories c ON p.categoryID = c.categoryID
// 		WHERE 
// 			p.categoryID = ?;
// 	`, id)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	//looping data
// 	for rows.Next() {
// 		var product_name string
// 		var price, stock int

// 		//check error
// 		err = rows.Scan(&product_name, &price, &stock)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}
// 		//print data
// 		fmt.Printf("Name: %s | Price: Rp %d | Available Stock: %d\n",
// 			product_name, price, stock)
// 	}

// 	return nil
// }

// func (h *HandlerImpl) ShowProductCategories() error {
// 	fmt.Println("masuk handler")
// 	rows, err := h.DB.Query(`
// 	SELECT
// 		categoryID AS id,
// 		name AS categories
// 	FROM
// 		product_categories;
// 	`)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("List of product category")
// 	//looping data
// 	for rows.Next() {
// 		var id int
// 		var categories string

// 		//check error
// 		err = rows.Scan(&id, &categories)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}
// 		//print data
// 		fmt.Printf("%v: %s\n",
// 			id, categories)
// 	}

// 	return nil
// }
