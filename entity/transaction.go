package entity

import "time"

type Transaction struct {
	trxID       int
	userID      int
	productID   int
	purchasedAt time.Time
}
