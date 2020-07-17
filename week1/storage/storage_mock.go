package storage

import (
	"errors"

	"github.com/sekarminati/todo2/model"
)

type Mock struct{}

func (o Mock) Create(customer model.Customer) error {
	return nil
}

func (o Mock) Detail(id int) (model.Customer, error) {
	if id == 1 {
		return model.Customer{ID: 2, Name: "mock name", Address: "mock address", Phone: "08345729357"}, nil
	}
	return model.Customer{}, errors.New("todo not found")
}

// func (o Mock) Delete(id int32) error {
// 	return nil
// }

func (o Mock) List() ([]model.Customer, error) {
	panic("not yet")
}
