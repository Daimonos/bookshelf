package cmd

import (
	"log"
	"strconv"

	"github.com/daimonos/go-bookshelf/data"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a book",
	Run:   AddCmd,
}

// AddCmd is the handler for the add cli command
func AddCmd(cmd *cobra.Command, args []string) {
	log.Println(args)
	var isRead, onLoan bool
	var err error
	if len(args) != 6 {
		log.Fatal("Expect 6 arguments: Title, Author, isRead, IsOnLoan, LoanedTo, ISBN")
	}
	isRead, err = strconv.ParseBool(args[2])
	onLoan, err = strconv.ParseBool(args[3])
	if err != nil {
		log.Fatalf("Error Parsing book from argument: %s", args[2])
	}
	book := data.Book{
		Title:    args[0],
		Author:   args[1],
		IsRead:   isRead,
		IsOnLoan: onLoan,
		LoanedTo: args[4],
		ISBN:     args[5],
	}
	_, err = store.AddBook(book)
	if err != nil {
		log.Fatalf("Error adding book: %s", err)
	}
	log.Println("Book Created!")

}
