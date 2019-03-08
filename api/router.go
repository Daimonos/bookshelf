package api

import (
	"github.com/daimonos/go-bookshelf/data"
	"github.com/gorilla/mux"
)

var store data.Store

func NewRouter(s *data.Store) *mux.Router {
	store = *s
	r := mux.NewRouter()
	r.HandleFunc("/list", HandleList)
	return r
}
