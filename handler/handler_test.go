package handler

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProductCategoryById(t *testing.T) {
	// Step 1: Create a mock database and mock object
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close() // Ensure the mockDB is closed after the test

	// Step 2: Set the expectation for the Exec query
	mock.ExpectExec("UPDATE product_categories SET name = \\? WHERE categoryID = \\?").
		WithArgs("New Category Name", 1).         // The values to replace the placeholders
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful update with 1 row affected

	// Step 3: Call the method you're testing
	handler := NewHandler(mockDB) // Assuming NewHandler initializes the handler
	err = handler.UpdateProductCategoryById(1, "New Category Name")

	// Step 4: Assertions
	// Check if there was no error during the execution
	assert.NoError(t, err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestBuyProduct(t *testing.T) {
	// Step 1: Create a mock database and mock object
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set the expectation for the Exec query
	mock.ExpectExec(`INSERT INTO transactions \(userID, productID, purchasedAt\) VALUES \(\?, \?, NOW\(\)\);`).
		WithArgs(1, 2).                           // Simulate the userID and productID being passed as arguments
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insert with 1 row affected

	// Step 3: Call the method you're testing
	handler := NewHandler(mockDB) // Assuming NewHandler initializes the handler
	err = handler.BuyProduct(1, 2)

	// Step 4: Assertions
	assert.NoError(t, err) // Check if no error occurred

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestShowOrders(t *testing.T) {
	// Step 1: Create a mock database and mock object
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set the expectation for the Query
	mock.ExpectQuery(`SELECT 
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
			MONTH\(t.purchasedAt\) = \? AND YEAR\(t.purchasedAt\) = \?
		ORDER BY 
			t.purchasedAt ASC;`).
		WithArgs(11, 2024). // Expected month and year (November 2024)
		WillReturnRows(sqlmock.NewRows([]string{
			"transaction_id", "user_name", "user_email", "product_name", "product_price", "purchase_date",
		}).
			AddRow(1, "John Doe", "johndoe@example.com", "Gaming Laptop", 1500000, "2024-11-05").
			AddRow(2, "Jane Smith", "janesmith@example.com", "Wireless Mouse", 500000, "2024-11-06"))

	// Step 3: Call the method you're testing
	handler := NewHandler(mockDB)      // Assuming NewHandler initializes the handler
	err = handler.ShowOrders(11, 2024) // November 2024

	// Step 4: Assertions
	assert.NoError(t, err) // Check if no error occurred

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestDeleteTransactionById(t *testing.T) {
	// Step 1: Create a mock database and mock object
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set the expectation for the DELETE query
	mock.ExpectExec("DELETE FROM transactions WHERE trxID = \\?").
		WithArgs(1).                              // The transaction ID to delete
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful delete with 1 row affected

	// Step 3: Call the method you're testing
	handler := NewHandler(mockDB)          // Assuming NewHandler initializes the handler
	err = handler.DeleteTransactionById(1) // Pass the transaction ID to delete

	// Step 4: Assertions
	assert.NoError(t, err) // Check if no error occurred

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestCreateNewUser(t *testing.T) {
	// Step 1: Bikin Mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set expectation
	mock.ExpectExec("INSERT INTO users \\(name, email\\) VALUES \\(\\?, \\?\\)").
		WithArgs("Dharma Satrya", "dharmasatrya@gmail.com"). // valuenya
		WillReturnResult(sqlmock.NewResult(1, 1))            // simulasi 1 row nambah

	// Step 3: panggil
	handler := NewHandler(mockDB)
	err = handler.CreateUser("Dharma Satrya", "dharmasatrya@gmail.com")

	// Step 4: assertion
	assert.NoError(t, err) // Check error

	// check expected output dari function udah sama dengan mock
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestAddProduct(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	mock.ExpectExec("INSERT INTO products \\(name, price, quantity, categoryID\\) VALUES \\(\\?, \\?, \\?, \\?\\)").
		WithArgs("The Sims 4", 100000, 10, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	handler := NewHandler(mockDB)
	err = handler.AddProduct("The Sims 4", 100000, 10, 2)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestCreateNewCategory(t *testing.T) {
	// Step 1: Bikin Mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set expectation
	mock.ExpectExec(`INSERT INTO product_categories \(name\) VALUES \(\?\)`).
		WithArgs("New Product Category").         // valuenya
		WillReturnResult(sqlmock.NewResult(1, 1)) // simulasi 1 row nambah

	// Step 3: panggil
	handler := NewHandler(mockDB)
	err = handler.CreateNewCategory("New Product Category")

	// Step 4: assertion
	assert.NoError(t, err) // Check error

	// check expected output dari function udah sama dengan mock
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}

func TestUpdateUserNameById(t *testing.T) {
	// Step 1: Bikin Mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()

	// Step 2: Set expectation
	mock.ExpectExec("UPDATE users SET name = \\? WHERE userID = \\?").
		WithArgs("New User Name", 1).             // The new name and user ID to update
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful update with 1 row affected

	// Step 3: panggil
	handler := NewHandler(mockDB)                        // Assuming NewHandler initializes the handler
	err = handler.UpdateUserNameById(1, "New User Name") // Pass the user ID and the new name

	// Step 4: assertion
	assert.NoError(t, err) // Check if no error occurred

	// check expected output dari function udah sama dengan mock
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expected query was not executed: %v", err)
	}
}
