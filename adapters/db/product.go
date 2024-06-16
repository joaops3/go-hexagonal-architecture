package db

import (
	"database/sql"

	"github.com/joaops3/go-exagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	Db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb{
	return &ProductDb{
		Db: db,
	}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error){
	var product application.Product

	stmt, err := p.Db.Prepare("select id, name, price status from products where id=?")

	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error){

	var rows int

	p.Db.QueryRow("select id from products where id=?").Scan(&rows)
	
	if rows != 0 {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}else {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	}
	
	return product, nil
}


func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error){

	stmt, err := p.Db.Prepare(`insert into products(id, name, price status) values(?,?,?,?)`)

	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}


func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error){

	
	_, err := p.Db.Exec(`update products set name = ?, price = ?, status = ? where id = ?`, product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}

	return product, nil
}