package data

import "time"

type Book struct {
	Title     string    `json:"title"`
	Author    string    `json:"isbn"`
	IsRead    bool      `json:"is_read"`
	IsOnLoan  bool      `json:"is_on_loan"`
	LoanedTo  string    `json:"loaned_to"`
	CreatedAt time.Time `json:"created_at"`
}
