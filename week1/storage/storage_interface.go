package storage

import "github.com/sekarminati/todo2/model"

type Storage interface {
	Create(customer model.Customer) error
	Detail(id int) (model.Customer, error)
	// Delete(id int32) error
	List() ([]model.Customer, error)
}
