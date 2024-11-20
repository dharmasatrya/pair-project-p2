package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Handler is the handler interface
type Handler interface {
	ShowUserSpending() error
	ShowOrders(month, year int) error
	ShowStocks() error
	SebuahFucntionCRUD() error
}

// HandlerImpl is the handler implementation
type HandlerImpl struct {
	DB *sql.DB
}

// NewHandler creates a new handler
func NewHandler(db *sql.DB) *HandlerImpl {
	return &HandlerImpl{
		DB: db,
	}
}

// FUNCTION BUAT NAMPILIN REPORT
func (h *HandlerImpl) ShowUserSpending() error {
	rows, err := h.DB.Query(`
SELECT 
    u.name AS user_name,
    SUM(p.price) AS total_spending
FROM 
    transactions t
JOIN 
    users u ON t.userID = u.userID
JOIN 
    products p ON t.productID = p.productID
GROUP BY 
    u.userID, u.name
ORDER BY 
    total_spending DESC;
	`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Total Game Sales Report:")
	//looping data
	for rows.Next() {
		var GAME_NAME string
		var TOTAL_SALES int

		//check error
		err = rows.Scan(&GAME_NAME, &TOTAL_SALES)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}
		//print data
		fmt.Printf("Game: %v, Total Sales: %v\n", GAME_NAME, TOTAL_SALES)
	}

	return nil
}

func (h *HandlerImpl) ShowOrders(month, year int) error {

	rows, err := h.DB.Query(`
		SELECT 
			t.trxID AS transaction_id,
			u.name AS user_name,
			u.email AS user_email,
			p.name AS product_name,
			p.price AS product_price,
			t.purchasedAt AS purchase_date
		FROM 
			transactions t
		JOIN 
			users u ON t.userID = u.userID
		JOIN 
			products p ON t.productID = p.productID
		WHERE 
			MONTH(t.purchasedAt) = ? AND YEAR(t.purchasedAt) = ?
		ORDER BY 
			t.purchasedAt ASC;
	`, month, year)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Total Game Sales Report:")

	if !rows.Next() {
		fmt.Println("No records found on this timerange")
	}

	for rows.Next() {
		var transactionID int
		var userName, userEmail, productName string
		var productPrice int
		var purchaseDate string

		err = rows.Scan(&transactionID, &userName, &userEmail, &productName, &productPrice, &purchaseDate)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		// Print each transaction
		fmt.Printf("Transaction ID: %d | User: %s (%s) | Product: %s | Price: %d | Date: %s\n",
			transactionID, userName, userEmail, productName, productPrice, purchaseDate)
	}

	// Check for errors during rows iteration
	if err = rows.Err(); err != nil {
		log.Print("Error during rows iteration: ", err)
		return err
	}

	return nil
}

func (h *HandlerImpl) ShowStocks() error {
	rows, err := h.DB.Query(`
		
	`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Total Game Sales Report:")
	//looping data
	for rows.Next() {
		var GAME_NAME string
		var TOTAL_SALES int

		//check error
		err = rows.Scan(&GAME_NAME, &TOTAL_SALES)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}
		//print data
		fmt.Printf("Game: %v, Total Sales: %v\n", GAME_NAME, TOTAL_SALES)
	}

	return nil
}

func (h *HandlerImpl) SebuahFucntionCRUD() error {
	rows, err := h.DB.Query(`
	QUERY CRUD
`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("crud selesai")

	return nil
}
