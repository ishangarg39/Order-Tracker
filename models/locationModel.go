package models

import (
	dbconnection "OrderTracker/dbConnection"
	"log"
)

type Location struct {
	// This struct will represent a location in the database.
	ID      int64
	Name    string `binding:"required"`
	Type    string `binding:"required"`
	Address string `binding:"required"`
}

// SaveLocationQuery saves a location to the database.
func (e Location) SaveLocationQuery() error {
	// This function will save a location to the database.
	// This function will be called from the controllers.go file.

	query := "INSERT INTO locations (name, type, address) VALUES ($1, $2, $3)"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Type, e.Address)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err

}

// GetLocationQuery gets a location from the database.
func GetLocationQuery() ([]Location, error) {
	// This function will get a location from the database.
	// This function will be called from the controllers.go file.

	query := "SELECT * FROM locations"

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

	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var location Location

		err := rows.Scan(&location.ID, &location.Name, &location.Type, &location.Address)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		locations = append(locations, location)
		log.Println(locations)
	}

	return locations, err

}

// UpdateLocationQuery updates a location in the database.
func (e Location) UpdateLocationQuery(id int64) error {
	// This function will update a location in the database.
	// This function will be called from the controllers.go file.

	query := "UPDATE locations SET name = $1, type = $2, address = $3 WHERE id = $4"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Type, e.Address, e.ID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err

}

// DeleteLocationQuery deletes a location from the database.
func DeleteLocationQuery(id int64) error {
	// This function will delete a location from the database.
	// This function will be called from the controllers.go file.

	query := "DELETE FROM locations WHERE id = $1"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

func GetLocationByIDQuery(id int64) (Location, error) {
	// This function will get a location from the database.
	// This function will be called from the controllers.go file.

	query := "SELECT * FROM locations WHERE id = $1"

	stmt, err := dbconnection.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return Location{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
		return Location{}, err
	}

	defer rows.Close()

	var location Location
	for rows.Next() {
		err := rows.Scan(&location.ID, &location.Name, &location.Type, &location.Address)

		if err != nil {
			log.Fatal(err)
			return Location{}, err
		}

	}

	return location, err

}
