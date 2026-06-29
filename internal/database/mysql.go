package database

import (
	"database/sql"
	"fmt"
	"log"

	"takah-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	cfg := config.LoadDatabaseConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db

	log.Println("✅ Database MySQL berhasil terhubung")
}