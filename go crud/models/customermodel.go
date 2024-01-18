package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/entities"
)

type customermodel struct {
	conn *sql.DB
}

func NewCustomerModel() *customermodel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &customermodel{
		conn: conn,
	}
}

func (p *customermodel) FindAll() ([]entities.Customer, error) {
	rows, err := p.conn.Query("select * from customer")
	if err != nil {
		return []entities.Customer{}, err
	}
	defer rows.Close()

	var dataCustomer []entities.Customer
	for rows.Next() {
		var customer entities.Customer
		rows.Scan(&customer.Id,
			&customer.NamaLengkap,
			&customer.NomorHp,
			&customer.Merk,
			&customer.Alamat,
			&customer.Masalah)

		dataCustomer = append(dataCustomer, customer)
	}

	return dataCustomer, nil
}

func (p *customermodel) Create(customer entities.Customer) bool {

	result, err := p.conn.Exec("insert into customer (nama_lengkap, no_hp, merk, alamat, masalah) values(?,?,?,?,?)",
		customer.NamaLengkap, customer.NomorHp, customer.Merk, customer.Alamat, customer.Masalah)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *customermodel) Find(id int64, customer *entities.Customer) error {

	return p.conn.QueryRow("select * from customer where id = ?", id).Scan(
		&customer.Id,
		&customer.NamaLengkap,
		&customer.NomorHp,
		&customer.Merk,
		&customer.Alamat,
		&customer.Masalah)

}

func (p *customermodel) Update(customer entities.Customer) error {

	_, err := p.conn.Exec(
		"update customer set nama_lengkap = ?, no_hp = ?, merk = ?, alamat = ?, masalah=? where id = ?",
		customer.NamaLengkap, customer.NomorHp, customer.Merk, customer.Alamat, customer.Masalah, customer.Id)

	if err != nil {
		return err
	}

	return nil

}

func (p *customermodel) Delete(id int64) {
	p.conn.Exec("delete from customer where id = ?", id)
}
