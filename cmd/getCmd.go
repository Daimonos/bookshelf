package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Book by Id",
	Run:   GetCmd,
}

func GetCmd(cmd *cobra.Command, args []string) {
	key, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Getting Book by ID: %d\n", key)
	book, err := store.GetBookByKey(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(book)
}
