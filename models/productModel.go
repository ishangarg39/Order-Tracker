package models

import (
	dbconnection "OrderTracker/dbConnection"
	"log"
)

type Product struct {
	// This struct will represent a product in the database.
	ID          int64
	Name        string  `binding:"required"`
	Description string  `binding:"required"`
	Price       float64 `binding:"required"`
	SupplierId  int64   `binding:"required"`
}

func GetProductQuery() ([]Product, error) {
	// This function will get a product from the database.
	// This function will be called from the controllers.go file.

	query := "SELECT * FROM products"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(1)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	productSlice := []Product{}

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		productSlice = append(productSlice, product)
	}

	return productSlice, nil
}

func (e Product) SaveProductQuery() error {
	// This function will save a product to the database.
	// This function will be called from the controllers.go file.

	query := "INSERT INTO products (name, description, price, supplier_id) VALUES ($1, $2, $3, $4)"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Price, e.SupplierId)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err

}
