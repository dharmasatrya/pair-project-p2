package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type GameRepository interface {
	UpdateProductCategoryById(id int, newName string) error
}

type RepoImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *RepoImpl {
	return &RepoImpl{
		DB: db,
	}
}

// FUNCTIONS BUAT CRUD
func (h *RepoImpl) UpdateProductCategoryById(id int, newName string) error {

	rows, err := h.DB.Query(`
	UPDATE product_categories
	SET name = ?
	WHERE categoryID = ?;
	`, newName, id)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

func (h *RepoImpl) UpdateUserNameById(id int, newName string) error {
	rows, err := h.DB.Query(`
	UPDATE users
	SET name = ?
	WHERE userID = ?;
	`, newName, id)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

func (h *RepoImpl) CreateNewCategory(name string) error {
	rows, err := h.DB.Query(`
	INSERT INTO product_categories (name)
	VALUES (?);`, name)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

func (h *RepoImpl) BuyProduct(userID, productID int) error {

	rows, err := h.DB.Query(`
	INSERT INTO transactions (userID, productID, purchasedAt)
	VALUES (?, ?, NOW());
	`, userID, productID)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

func (h *RepoImpl) DeleteTransactionById(transactionID int) error {

	rows, err := h.DB.Query(`
		DELETE FROM transactions WHERE trxID = ?;
	`, transactionID)
	if err != nil {
		log.Print("Error deleting records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

// FUNCTION BUAT NAMPILIN REPORT
func (h *RepoImpl) ShowUserSpending() error {
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

	return rows
}

func (h *RepoImpl) ShowOrders(month, year int) error {

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

	return rows
}

func (h *RepoImpl) ShowStocks() error {
	rows, err := h.DB.Query(`
		SELECT 
			pc.name AS category_name,
			p.name AS product_name,
			p.quantity AS current_quantity
		FROM 
			products p
		JOIN 
			product_categories pc ON p.categoryID = pc.categoryID
		ORDER BY 
			pc.name, p.name;
	`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return rows
}

func (h *RepoImpl) ShowProductByCategoryId(id int) error {
	rows, err := h.DB.Query(`
		SELECT 
			p.name AS product_name,
			p.price AS price,
			p.quantity AS stock
		FROM 
			products p
		JOIN 
			product_categories c ON p.categoryID = c.categoryID
		WHERE 
			p.categoryID = ?;
	`, id)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}

func (h *RepoImpl) ShowProductCategories() error {
	rows, err := h.DB.Query(`
	SELECT
		categoryID AS id,
		name AS categories
	FROM
		product_categories;
	`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	return nil
}
