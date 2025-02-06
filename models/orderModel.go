package models

import dbconnection "OrderTracker/dbConnection"

type Order struct {
	ID         int64  `json:"id"`
	Date       string `json:"date"`
	CustomerId int64  `json:"customerId"`
	SupplierId int64  `json:"supplierId"`
	ProductId  int64  `json:"productId"`
}

func GetOrderQuery() ([]Order, error) {
	// This function will return a slice of orders from the database
	// This function will be called from the controllers.

	// Get the orders here.
	query := "SELECT * FROM orders"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	orderSlice := []Order{}

	for rows.Next() {
		var order Order

		err := rows.Scan(&order.ID, &order.Date, &order.CustomerId, &order.SupplierId, &order.ProductId)

		if err != nil {
			return nil, err
		}

		orderSlice = append(orderSlice, order)
	}

	return orderSlice, nil

}

func (o Order) SaveOrderQuery() error {
	// This function will save an order to the database
	// This function will be called from the controllers.

	// Save the order here.
	query := "INSERT INTO orders (date, customer_id, supplier_id, product_id) VALUES ($1, $2, $3, $4)"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(o.Date, o.CustomerId, o.SupplierId, o.ProductId)

	if err != nil {
		return err
	}

	return nil
}
