package transaction

import "time"

// Transaction é o modelo para a transação financeira
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

// Transactions é mma coleção de transações
type Transactions []Transaction
