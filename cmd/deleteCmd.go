package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Book",
	Run:   DeleteCmd,
}

func DeleteCmd(cmd *cobra.Command, args []string) {
	key, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	err = store.DeleteBookByKey(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted book with key: %s\n", key)
}
