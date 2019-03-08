package data

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type Store struct {
	db *bolt.DB
}

func (s *Store) Init() {
	log.Println("initializing store")
	var err error
	s.db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// create buckets
	err = s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("books"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) AddBook(title, author string, isRead, isOnLoan bool, onLoanTo string) (Book, error) {
	var book Book
	err := s.db.Update(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("books"))
		book := Book{
			Title:     title,
			Author:    author,
			IsRead:    isRead,
			IsOnLoan:  isOnLoan,
			LoanedTo:  onLoanTo,
			CreatedAt: time.Now(),
		}
		if !book.IsOnLoan {
			book.LoanedTo = ""
		}
		var buffer []byte
		buffer, err = GetBufferFromStruct(book)
		if err != nil {
			return err
		}
		err = b.Put([]byte(book.Title), buffer)
		return nil
	})
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (s *Store) GetAllBooks() ([]Book, error) {
	var books []Book
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		err := b.ForEach(func(k, v []byte) error {
			var book Book
			var err error
			err = json.Unmarshal(v, &book)
			if err != nil {
				return err
			}
			books = append(books, book)
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBufferFromStruct(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
