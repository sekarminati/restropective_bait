package main

import (
	"fmt"
	"log"

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
		ID:      1,
		Name:    "Sekar",
		Address: "Jakarta",
		Phone:   "085822339848",
	}
	if err := memStore.Create(obj); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data berhasil diinput!\n\n")

	// DETAIL
	cus, err := memStore.Detail(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID = %v, Nama = %v, Alamat = %v, No. Telp = %v\n\n", cus.ID, cus.Name, cus.Address, cus.Phone)

	// SELECT ALL/LIST
	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DATA PELANGGAN")
	for i, cus := range list {
		fmt.Printf("Row %v: ID = %v, Nama = %v, Alamat = %v, No. Telp = %v\n", i, cus.ID, cus.Name, cus.Address, cus.Phone)
	}
}
