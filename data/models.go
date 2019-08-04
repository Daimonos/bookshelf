package data

// Book struct
type Book struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	IsRead   bool   `json:"is_read"`
	IsOnLoan bool   `json:"is_on_loan"`
	LoanedTo string `json:"loaned_to"`
	ISBN     string `json:"isbn"`
}
