package models

import dbconnection "OrderTracker/dbConnection"

type Supplier struct {
	// This struct will represent a supplier in the database.
	ID         int64
	Name       string `binding:"required"`
	Contact    string `binding:"required"`
	LocationID int64  `binding:"required"`
}

func GetSupplierQuery() ([]Supplier, error) {
	// This function will return a slice of suppliers from the database
	// This function will be called from the controllers.
	// Get the suppliers here.
	query := "SELECT * FROM suppliers"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	supplierSlice := []Supplier{}

	for rows.Next() {
		var supplier Supplier

		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Contact, &supplier.LocationID)

		if err != nil {
			return nil, err
		}

		supplierSlice = append(supplierSlice, supplier)
	}

	return supplierSlice, nil

}

func (s Supplier) SaveSupplierQuery() error {
	// This function will save a supplier to the database
	// This function will be called from the controllers.
	// Save the supplier here.
	query := "INSERT INTO suppliers (name, contact, location_id) VALUES ($1, $2, $3)"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(s.Name, s.Contact, s.LocationID)

	if err != nil {
		return err
	}

	return nil

}
