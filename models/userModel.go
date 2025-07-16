package models

import (
	"database/sql"
	"trial-go/config"
)

type Users struct {
	ID       int
	Username string
	Name     string
	Email    *string
	Phone    *string
	Level    string
}

func GetAllUsers() ([]*Users, error) {
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
	var users []*Users

	for rows.Next() {
		user := &Users{}
		var email sql.NullString
		var phone sql.NullString

		err := rows.Scan(&user.ID, &user.Username, &user.Name, &email, &phone, &user.Level)
		if err != nil {
			return nil, err
		}

		// handling ketika data nullable
		if email.Valid {
			user.Email = &email.String
		}

		if phone.Valid {
			user.Phone = &phone.String
		}
		users = append(users, user)
	}

	return users, nil
}
