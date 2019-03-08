package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/daimonos/go-bookshelf/api"
	"github.com/daimonos/go-bookshelf/data"
	"github.com/spf13/cobra"
)

var store data.Store

var rootCmd = &cobra.Command{
	Use: "bookshelf",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a book",
	Run:   AddCmd,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Books",
	Run:   ListCmd,
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start API",
	Run:   StartApi,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCmd(cmd *cobra.Command, args []string) {
	log.Println(args)
	var isRead, onLoan bool
	var err error
	if len(args) != 5 {
		log.Fatal("Expect 5 arguments: Title, Author, isRead, IsOnLoan, LoanedTo")
	}
	isRead, err = strconv.ParseBool(args[2])
	onLoan, err = strconv.ParseBool(args[3])
	if err != nil {
		log.Fatalf("Error Parsing book from argument: %s", args[2])
	}
	_, err = store.AddBook(args[0], args[1], isRead, onLoan, args[4])
	if err != nil {
		log.Fatalf("Error adding book: %s", err)
	}
	log.Println("Book Created!")

}

func ListCmd(cmd *cobra.Command, args []string) {
	books, err := store.GetAllBooks()
	if err != nil {
		log.Fatalf("Error getting all books: %s", err)
	}
	fmt.Printf("%-75s %-20s %-10s %-10s %-15s\n", "Title (Key)", "Author", "Read", "Loan", "Loaned To")
	for _, b := range books {
		fmt.Printf("%-75s %-20s %-10s %-10s %-15s\n", b.Title, b.Author, strconv.FormatBool(b.IsRead), strconv.FormatBool(b.IsOnLoan), b.LoanedTo)
	}
	fmt.Println("---\nDone")
}

func StartApi(cmd *cobra.Command, args []string) {
	r := api.NewRouter(&store)
	http.ListenAndServe(":8080", r)
}

func init() {
	store = data.Store{}
	store.Init()
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(apiCmd)
}
