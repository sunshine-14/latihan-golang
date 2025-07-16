package models

import (
	"database/sql"
	"trial-go/config"
)

type Users struct {
	ID       int
	Username string
	Name     string
	Email    string
	Phone    string
	Level    string
}

func GetAllUsers() ([]Users, error) {
	// panggil db config
	db, err := config.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	// query kan disini
	rows, err := db.Query("SELECT id, username, nama, email, phone, level FROM users")
	if err != nil {
		return nil, err
	}

	// tutup koneksi ke DB
	defer rows.Close()

	// variable yg akan direturn
	var user []Users

	for rows.Next() {
		var data Users
		var email sql.NullString
		var phone sql.NullString

		err := rows.Scan(&data.ID, &data.Username, &data.Name, &email, &phone, &data.Level)
		if err != nil {
			return nil, err
		}

		// handling ketika data nullable
		if email.Valid {
			data.Email = email.String
		}

		if phone.Valid {
			data.Phone = phone.String
		}
		user = append(user, data)
	}

	return user, nil
}
