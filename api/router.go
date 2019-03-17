package api

import (
	"github.com/daimonos/go-bookshelf/data"
	"github.com/gorilla/mux"
)

var store data.Store

func NewRouter(s *data.Store) *mux.Router {
	store = *s
	r := mux.NewRouter()
	r.HandleFunc("/books", HandleBookList).Methods("GET")
	r.HandleFunc("/books", HandleBookPost).Methods("POST")
	r.HandleFunc("/books/{id}", HandleGetByKey).Methods("GET")
	r.HandleFunc("/books/{id}", HandleDeleteByKey).Methods("DELETE")
	return r
}
