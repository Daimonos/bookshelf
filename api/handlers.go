package api

import (
	"encoding/json"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	books, err := store.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	var data []byte
	data, err = json.Marshal(books)
	w.Write(data)

}
