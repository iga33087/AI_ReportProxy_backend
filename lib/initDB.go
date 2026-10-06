package lib

import (
	//"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/test.db")
	if err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// 建立資料表
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS devices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		productId TEXT NOT NULL,
		name TEXT NOT NULL,
		key TEXT NOT NULL,
		uuid TEXT NOT NULL UNIQUE,
		cpu_usage_summary TEXT,
		memory_usage_summary TEXT,
		new_connection_summary TEXT,
		active_connection_summary TEXT,
		online_verification_summary TEXT,
		online_users_summary TEXT,
		top20UserTrafficRanking TEXT,
		top20UserTrafficGroupRanking TEXT,
		top20ServiceTrafficRanking TEXT,
		top20ServiceTrafficTypeRanking TEXT,
		top20DomainTrafficRanking TEXT,
		top20DomainTrafficTypeRanking TEXT
	);

	CREATE TABLE IF NOT EXISTS reports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		deviceId INTEGER NOT NULL,
		reportType INTEGER NOT NULL,
		reportData TEXT
	);
	`

	_, err := db.Exec(query)
	return err
}