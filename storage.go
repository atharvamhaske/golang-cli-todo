package main

import (
	"encoding/json"
	"os"
)

type Storage[t any] struct {
	FileName string
}

//below is the constructor function called as generic function

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName} //T here is placeholder and *Storage[T] is what function promises to give back

}

func (s *Storage[T]) save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}


func (s *Storage[T]) Load(data *T) error {
    fileData, err := os.ReadFile(s.FileName)
   
    if err != nil {
        return err
    }

    return json.Unmarshal(fileData, data)
}