package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sekarminati/todo2/model"
)

type Database struct{}

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (o Database) Create(customer model.Customer) error {
	db := InitDB()
	_, err := db.Exec("INSERT INTO customers (id, name, address, phone) VALUES"+
		"($1, $2, $3, $4);",
		customer.ID,
		customer.Name,
		customer.Address,
		customer.Phone)
	return err
}

func (o Database) List() ([]model.Customer, error) {
	db := InitDB()
	rows, err := db.Query("SELECT id, name, address, phone FROM customers LIMIT 10;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]model.Customer, 0)
	for rows.Next() {
		var cus model.Customer
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone)
		if err != nil {
			return nil, err
		}
		res = append(res, cus)
	}
	return res, nil
}

func (o Database) Detail(id int) (model.Customer, error) {
	db := InitDB()
	var cus model.Customer
	err := db.QueryRow("SELECT id, name, address, phone FROM customers WHERE id=$1;", id).Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone)
	if err != nil {
		log.Fatal(err)
	}
	return cus, nil
}
