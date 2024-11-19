package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Handler is the handler interface
type Handler interface {
	SebuahFucntionReport() error
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
func (h *HandlerImpl) SebuahFucntionReport() error {
	rows, err := h.DB.Query(`
		ISI SAMA QUERY REPORT
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
