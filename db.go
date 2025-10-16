package main

import (
	"encoding/json"
	"os"
)

type DB[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) *DB[T] {
	return &DB[T]{
		Filename: filename,
	}

}

func (db *DB[T]) Save(data T) error {

	// Convert data to JSON
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// Write JSON to file
	err = os.WriteFile(db.Filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB[T]) Load(data *T) error {
	fileData, err := os.ReadFile(db.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}
	return nil
}
