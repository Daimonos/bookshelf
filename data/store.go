package data

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// Store for holding the db
type Store struct {
	db *bolt.DB
}

// Init - Initializes the store
func (s *Store) Init() {
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

// AddBook - adds a book to the store
func (s *Store) AddBook(book Book) (Book, error) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("books"))
		if !book.IsOnLoan {
			book.LoanedTo = ""
		}
		book.ID, err = b.NextSequence()
		if err != nil {
			return err
		}
		var buffer []byte
		buffer, err = GetBufferFromStruct(book)
		if err != nil {
			return err
		}
		err = b.Put(itob(book.ID), buffer)
		return nil
	})
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

// GetBookByKey - Gets book details by key
func (s *Store) GetBookByKey(key uint64) (Book, error) {
	var book Book
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		v := b.Get(itob(key))
		json.Unmarshal(v, &book)
		return nil
	})
	if err != nil {
		return Book{}, err
	}
	if book.ID != key {
		return Book{}, errors.New("not found")
	}
	return book, nil
}

// DeleteBookByKey - Deletes a book by it's id
func (s *Store) DeleteBookByKey(key uint64) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		err := b.Delete(itob(key))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// GetAllBooks - returns all books in the store
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

func (s *Store) UpdateBook(key uint64, book Book) (Book, error) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		bytes, err := GetBufferFromStruct(book)
		if err != nil {
			return err
		}
		b.Put(itob(key), bytes)
		return nil
	})
	return book, err
}

// GetBufferFromStruct converts any struct into a byte array
func GetBufferFromStruct(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// itob converts a unint64 to a byte array
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
