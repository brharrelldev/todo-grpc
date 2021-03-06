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
	Id       string `json:"id"`
	Done     bool   `json:"done"`
}

type AllTodoData []TodoData

func (t AllTodoData) Iterate() <-chan TodoData {

	todo := make(chan TodoData)
	go func() {
		for _, entry := range t {
			select {
			case <-todo:
				close(todo)
				return
			case todo <- entry:

			}
		}
		close(todo)
	}()

	return todo

}

func NewDatabase() (*TodoDatabase, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/todo_db"))
	if err != nil {
		return nil, fmt.Errorf("error opening todoDB %v", err)
	}

	return &TodoDatabase{db: db}, nil
}

func (todo *TodoDatabase) Add(data *TodoData) (string, error) {

	var id string

	if err := todo.db.Update(func(txn *badger.Txn) error {
		todoId := uuid.NewV4()
		id = todoId.String()
		data.Id = id


		payload, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("error marshaling data %v", err)
		}


		entry := badger.NewEntry(todoId.Bytes(), payload)

		if err := txn.SetEntry(entry); err != nil {
			return fmt.Errorf("error writing new entry to DB %v", err)
		}

		return nil
	}); err != nil {
		return "", fmt.Errorf("error updating database %v", err)
	}

	return id, nil
}

func (todo *TodoDatabase) Get(id string) (*TodoData, error) {

	var td *TodoData
	if err := todo.db.View(func(txn *badger.Txn) error {

		val, err := txn.Get([]byte(id))
		if err != nil {
			return fmt.Errorf("error getting id %v", err)
		}

		byteVal, err := val.ValueCopy(nil)
		if err != nil {
			return fmt.Errorf("error copying value %v", err)
		}

		if err := json.NewDecoder(bytes.NewBuffer(byteVal)).Decode(&td); err != nil {
			return fmt.Errorf("error decoding to new json %v", err)
		}

		return nil

	}); err != nil {
		return nil, fmt.Errorf("error retrieving entry %v", err)
	}

	if td == nil {
		return nil, errors.New("value not found, or serialization error")
	}

	return td, nil

}

func (todo *TodoDatabase) GetAll() (AllTodoData, error) {
	var todos AllTodoData
	if err := todo.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)

		defer iter.Close()

		for iter.Rewind(); iter.Valid(); iter.Next() {
			entry := iter.Item()
			val, err := entry.ValueCopy(nil)
			if err != nil {
				return fmt.Errorf("error copying value %v", err)
			}

			fmt.Println(string(val))

			t := TodoData{}
			if err := json.NewDecoder(bytes.NewBuffer(val)).Decode(&t); err != nil {
				return fmt.Errorf("error creating new decoder %v", err)
			}

			todos = append(todos, t)
		}

		return nil

	}); err != nil {
		return nil, fmt.Errorf("error occured when retrieving entries %v", err)
	}

	return todos, nil
}

func (todo *TodoDatabase) DeleteTodo(id string) error {

	if err := todo.db.Update(func(txn *badger.Txn) error {

		if err := txn.Delete([]byte(id)); err != nil{
			return fmt.Errorf("error deleting id %v", err)
		}

		return nil

	}); err != nil{
		return fmt.Errorf("error deleting transaction %v", err)
	}

	return nil

}
