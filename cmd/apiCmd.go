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
	Run:   StartApi,
}

func StartApi(cmd *cobra.Command, args []string) {
	r := api.NewRouter(&store)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
