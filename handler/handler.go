package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Handler is the handler interface
type Handler interface {
	ShowTotalGameSales() error
	ShowMostPopularGames() error
	ShowTotalGamesRevenue() error
	ShowTotalPlayer() error
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

func (h *HandlerImpl) ShowTotalGameSales() error {
	rows, err := h.DB.Query(`SELECT 
    games.name AS GAME_NAME, 
    COUNT(transactions.trxID) AS TOTAL_SALES
FROM 
    games
LEFT JOIN 
    transactions ON games.gameID = transactions.gameID
GROUP BY 
    games.gameID, games.name
ORDER BY 
    total_sales DESC;`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Total Game Sales Report:")
	for rows.Next() {
		var GAME_NAME string
		var TOTAL_SALES int
		err = rows.Scan(&GAME_NAME, &TOTAL_SALES)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("Game: %v, Total Sales: %v\n", GAME_NAME, TOTAL_SALES)
	}

	return nil
}

func (h *HandlerImpl) ShowMostPopularGames() error {
	rows, err := h.DB.Query(`SELECT 
    g.name AS game_name, 
    COUNT(DISTINCT t.playerID) AS total_unique_players
FROM 
    games g
LEFT JOIN 
    transactions t ON g.gameID = t.gameID
GROUP BY 
    g.gameID, g.name
ORDER BY 
    total_unique_players DESC;
`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Most Popular Game Report")
	for rows.Next() {
		var game_name string
		var total_unique_players int
		err = rows.Scan(&game_name, &total_unique_players)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("Game: %v, Unique Players: %v\n", game_name, total_unique_players)
	}

	return nil
}

func (h *HandlerImpl) ShowTotalGamesRevenue() error {
	rows, err := h.DB.Query(`SELECT 
    g.name AS game_name, 
    SUM(g.price) AS total_revenue
FROM 
    games g
JOIN 
    transactions t ON g.gameID = t.gameID
GROUP BY 
    g.name
ORDER BY 
    total_revenue DESC;
`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Total Revenue Per Game Report:")
	for rows.Next() {
		var game_name string
		var total_revenue float32
		err = rows.Scan(&game_name, &total_revenue)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("Game: %v, Total Revenue: $%v\n", game_name, total_revenue)
	}

	return nil
}

func (h *HandlerImpl) ShowTotalPlayer() error {
	rows, err := h.DB.Query(`
	SELECT 
    g.name AS game_name, 
    COUNT(DISTINCT t.playerID) AS total_unique_players
FROM 
    games g
LEFT JOIN 
    transactions t ON g.gameID = t.gameID
GROUP BY 
    g.gameID, g.name
ORDER BY 
    total_unique_players DESC;
`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("Player Count Per Game Report")
	for rows.Next() {
		var game_name string
		var total_unique_players int
		err = rows.Scan(&game_name, &total_unique_players)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("Game: %v, Player Count: %v\n", game_name, total_unique_players)
	}

	return nil
}
