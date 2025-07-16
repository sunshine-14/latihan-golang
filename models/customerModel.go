package models

import (
	"trial-go/config"
)

type Customer struct {
	ID            int
	Customer_ID   string
	Customer_Name string
}

func GetAllCustomers() ([]*Customer, error) {
	// start to connect to MySQL database
	db, err := config.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, customer_id, customer_name FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*Customer

	for rows.Next() {
		cust := &Customer{}
		// Scan data ke struct
		err := rows.Scan(&cust.ID, &cust.Customer_ID, &cust.Customer_Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, cust)
	}

	return customers, nil
}
