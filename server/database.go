package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger"
	uuid "github.com/satori/go.uuid"
)


type TodoDatabase struct {
	db *badger.DB
}

type TodoData struct {
	TaskName string `json:"task_name"`
	Done bool `json:"done"`
}

func NewDatabase() (*TodoDatabase, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp"))
	if err != nil{
		return nil, fmt.Errorf("error opening todoDB %v", err)
	}

	return &TodoDatabase{db: db}, nil
}

func (todo *TodoDatabase) Add(data *TodoData) (string,error){

	var id string

	if err := todo.db.Update(func(txn *badger.Txn) error {

		payload, err := json.Marshal(data)
		if err != nil{
			return fmt.Errorf("error marshaling data")
		}

		todoId := uuid.NewV4()
		id = todoId.String()

		entry := badger.NewEntry(todoId.Bytes(), payload)

		if err := txn.SetEntry(entry); err != nil{
			return fmt.Errorf("error writing new entry to DB %v", err)
		}


		if err := txn.Commit(); err != nil{
			return fmt.Errorf("transaction could not be commited %v", err)
		}



		return nil
	}); err != nil{
		return "", fmt.Errorf("error updating database %v", err)
	}

	return id, nil
}

func (todo *TodoDatabase) Get(id string) (*TodoData, error)  {

	var td *TodoData
	if err := todo.db.View(func(txn *badger.Txn) error {

		val, err := txn.Get([]byte(id))
		if err != nil{
			return fmt.Errorf("error getting id %v", err)
		}

		byteVal, err := val.ValueCopy(nil)
		if err != nil{
			return fmt.Errorf("error copying value %v", err)
		}

		if err := json.NewDecoder(bytes.NewBuffer(byteVal)).Decode(&td); err != nil{
			return fmt.Errorf("error decoding to new json %v", err)
		}

		return nil

	}); err != nil{
		return nil, fmt.Errorf("error retrieving entry %v", err)
	}

	if td == nil{
		return nil, errors.New("value not found, or serialization error")
	}

	return td, nil

}
