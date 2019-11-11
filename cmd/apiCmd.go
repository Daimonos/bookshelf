package cmd

import (
	"log"
	"net/http"

	"github.com/daimonos/go-bookshelf/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start API",
	Run:   StartAPI,
}

// StartAPI is a command that will start up a REST api
func StartAPI(cmd *cobra.Command, args []string) {
	r := api.NewRouter(&store)
	log.Println("Listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
