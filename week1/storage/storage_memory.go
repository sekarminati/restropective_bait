package storage

import (
	"errors"

	"github.com/sekarminati/todo2/model"
)

type Memory struct {
	store map[int]model.Customer
}

func NewMemory() Memory {
	return Memory{
		store: make(map[int]model.Customer),
	}
}

func (o Memory) Create(customer model.Customer) error {
	o.store[customer.ID] = customer
	return nil
}

func (o Memory) Detail(id int) (model.Customer, error) {
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Customer{}, errors.New("customer not found")
}

// func (o Memory) Delete(id int32) error {
// 	delete(o.store, id)
// 	return nil
// }

func (o Memory) List() ([]model.Customer, error) {
	panic("not yet")
}
