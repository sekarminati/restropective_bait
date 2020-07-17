package main

import (
	"fmt"
	"log"

	"github.com/sekarminati/todo2/iterator"
	"github.com/sekarminati/todo2/model"

	"github.com/sekarminati/todo2/storage"
)

type StorageType int

const (
	StorageTypeMemory StorageType = iota
	StorageTypeMock
	StorageTypeDatabase
)

// Factory
func getStorage(t StorageType) storage.Storage {
	var s storage.Storage
	switch t {
	case StorageTypeMemory:
		s = storage.NewMemory()
	case StorageTypeMock:
		s = storage.Mock{}
	case StorageTypeDatabase:
		s = storage.Database{}
	default:
		log.Fatal("storage unknown")
	}
	return s
}

func main() {
	//inisialisasi
	var memStore = getStorage(StorageTypeDatabase)

	// CREATE
	obj := model.Customer{
		ID:      3,
		Name:    "Sekar",
		Address: "Jakarta",
		Phone:   "085822339848",
	}
	if err := memStore.Create(obj); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data berhasil diinput!\n\n")

	// DETAIL
	cus, err := memStore.Detail(obj.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID = %v, Nama = %v, Alamat = %v, No. Telp = %v\n\n", cus.ID, cus.Name, cus.Address, cus.Phone)

	// SELECT ALL/LIST
	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}

	// ITERATOR
	fmt.Println("DATA PELANGGAN")
	itr := iterator.CustCollection{Cust: list}.CreateCustCollection()

	for itr.HasNext() {
		v := itr.GetNext()
		fmt.Printf("ID = %v, Nama = %v, Alamat = %v, No. Telp = %v\n", v.ID, v.Name, v.Address, v.Phone)
	}

}
