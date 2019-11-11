package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Books",
	Run:   ListCmd,
}

// ListCmd is the list command handler for the cli
func ListCmd(cmd *cobra.Command, args []string) {
	books, err := store.GetAllBooks()
	if err != nil {
		log.Fatalf("Error getting all books: %s", err)
	}
	fmt.Printf("%-5s %-75s %-20s %-10s %-10s %-15s\n", "ID", "Title", "Author", "Read", "Loan", "Loaned To")
	for _, b := range books {
		fmt.Printf("%-5d %-75s %-20s %-10s %-10s %-15s\n", b.ID, b.Title, b.Author, strconv.FormatBool(b.IsRead), strconv.FormatBool(b.IsOnLoan), b.LoanedTo)
	}
	fmt.Println("---\nDone")
}
