package iterator

import (
	"github.com/sekarminati/todo2/model"
)

type Iterator interface {
	HasNext() bool
	GetNext() model.Customer
}

type CustomerIterator struct {
	index    int
	custList []model.Customer
}

func (o *CustomerIterator) HasNext() bool {
	if o.index < len(o.custList) {
		return true
	}
	return false
}

func (o *CustomerIterator) GetNext() model.Customer {
	if o.HasNext() {
		cust := o.custList[o.index]
		o.index++
		return cust
	}
	return model.Customer{}
}

type CustCollection struct {
	Cust []model.Customer
}

func (o CustCollection) CreateCustCollection() Iterator {
	return &CustomerIterator{
		index:    0,
		custList: o.Cust,
	}
}
