package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(ip:port)/database")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Koneksi ke MySQL berhasil!")
	return db, nil
}
