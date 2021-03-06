package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/daimonos/go-bookshelf/data"
	"github.com/gorilla/mux"
)

func HandleBookList(w http.ResponseWriter, r *http.Request) {
	books, err := store.GetAllBooks()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, books)
}

func HandleBookPost(w http.ResponseWriter, r *http.Request) {
	var book data.Book
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&book)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}
	book, err = store.AddBook(book)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusCreated, book)
}

func HandleGetByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	var id uint64
	var book data.Book
	id, err = strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
	}
	book, err = store.GetBookByKey(id)
	if err != nil {
		if err.Error() == "not found" {
			WriteError(w, http.StatusNotFound, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, book)
}

func HandleDeleteByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
	}
	err = store.DeleteBookByKey(id)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, id)
}

func HandleBookUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var book data.Book
	var err error
	reader := json.NewDecoder(r.Body)
	err = reader.Decode(&book)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
	}

	book, err = store.UpdateBook(id, book)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, book)

}

func WriteError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	body, err := json.Marshal(payload)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}
