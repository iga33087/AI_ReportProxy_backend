package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Device struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	UUID string `json:"uuid"`
	Key string `json:"key"`
	HardwareData HardwareData `json:"hardwareData"`
	//HardwareData HardwareData
	//UserTrafficData UserTrafficData
	//ServiceTrafficData ServiceTrafficData
	//DomainTrafficData DomainTrafficData
}

type HardwareData struct {
    CPU_Usage_Summary string `json:"cpuUsageSummary"`
	Memory_Usage_Summary string `json:"memoryUsageSummary"`
	New_Connection_Summary string `json:"newConnectionSummary"`
	Active_Connection_Summary string `json:"activeConnectionSummary"`
	Online_Verification_Summary string `json:"onlineVerificationSummary"`
	Online_Users_Summary string `json:"onlineUsersSummary"`
}

type UserTrafficData struct {
    Top20Ranking string
    Top20GroupRanking string
}

type ServiceTrafficData struct {
    Top20Ranking string
    Top20TypeRanking string
}

type DomainTrafficData struct {
    Top20Ranking string
    Top20TypeRanking string
}

// 初始化資料庫
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
		name TEXT NOT NULL,
		key TEXT NOT NULL,
		uuid TEXT NOT NULL UNIQUE
	);
	`

	_, err := db.Exec(query)
	return err
}

// 查詢全部
func GetDevice(db *sql.DB) ([]Device, error) {
	query := `SELECT id, name, uuid, key FROM devices`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	

	var devices []Device

	for rows.Next() {
		var device Device

		err := rows.Scan(
			&device.ID,
			&device.Name,
			&device.UUID,
			&device.Key,
		)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

// 新增
func CreateDevice(db *sql.DB, device Device) (int64, error) {
	query := `INSERT INTO devices (name, uuid, key)VALUES (?, ?, ?)`

	result, err := db.Exec(query, device.Name, device.UUID, device.Key)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// 修改
func UpdateDevice(db *sql.DB,device Device) (any, error) {
	query := `UPDATE devices SET name = ?, uuid = ?, key = ? WHERE id = ?`

	result, err := db.Exec(query, device.Name, device.UUID, device.Key, device.ID)
	if err != nil {
		return nil,err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil,err
	}

	return affected,nil
}

// 刪除
func DeleteDevice(db *sql.DB, id string) (any, error) {
	query := `
		DELETE FROM devices
		WHERE id = ?
	`

	result, err := db.Exec(query, id)
	if err != nil {
		return nil,err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil,err
	}

	return affected,nil
}