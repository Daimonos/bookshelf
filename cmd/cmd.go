package cmd

import (
	"fmt"
	"os"

	"github.com/daimonos/go-bookshelf/data"
)

var store data.Store

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	store = data.Store{}
	store.Init()
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(getCmd)
}
