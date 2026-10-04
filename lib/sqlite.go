package lib

import (
	//"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Device struct {
	ID    int `json:"id"`
	ProductId string `json:"productId"`
	Name  string `json:"name"`
	UUID string `json:"uuid"`
	Key string `json:"key"`
	HardwareData HardwareData `json:"hardwareData"`
	UserTrafficData UserTrafficData `json:"userTrafficData"`
	ServiceTrafficData ServiceTrafficData `json:"serviceTrafficData"`
	DomainTrafficData DomainTrafficData `json:"domainTrafficData"`
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
    Top20_UserTraffic_Ranking string `json:"top20UserTrafficRanking"`
    Top20_UserTraffic_Group_Ranking string `json:"top20UserTrafficGroupRanking"`
}

type ServiceTrafficData struct {
    Top20_ServiceTraffic_Ranking string `json:"top20ServiceTrafficRanking"`
    Top20_ServiceTraffic_Type_Ranking string `json:"top20ServiceTrafficTypeRanking"`
}

type DomainTrafficData struct {
    Top20_DomainTraffic_Ranking string `json:"top20DomainTrafficRanking"`
    Top20_DomainTraffic_Type_Ranking string `json:"top20DomainTrafficTypeRanking"`
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
	`

	_, err := db.Exec(query)
	return err
}

// 查詢全部
func GetDevice(db *sql.DB) ([]Device, error) {
	query := `SELECT 
	id, 
	productId,
	name, 
	uuid, 
	key, 
	cpu_usage_summary,
	memory_usage_summary,
	new_connection_summary,
	active_connection_summary,
	online_verification_summary,
	online_users_summary,
	top20UserTrafficRanking,
	top20UserTrafficGroupRanking,
	top20ServiceTrafficRanking,
	top20ServiceTrafficTypeRanking,
	top20DomainTrafficRanking,
	top20DomainTrafficTypeRanking 
	FROM devices`

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
			&device.ProductId,
			&device.Name,
			&device.UUID,
			&device.Key,
		    &device.HardwareData.CPU_Usage_Summary,
		    &device.HardwareData.Memory_Usage_Summary,
		    &device.HardwareData.New_Connection_Summary,
		    &device.HardwareData.Active_Connection_Summary,
		    &device.HardwareData.Online_Verification_Summary,
		    &device.HardwareData.Online_Users_Summary,
			&device.UserTrafficData.Top20_UserTraffic_Ranking,
			&device.UserTrafficData.Top20_UserTraffic_Group_Ranking,
			&device.ServiceTrafficData.Top20_ServiceTraffic_Ranking,
			&device.ServiceTrafficData.Top20_ServiceTraffic_Type_Ranking,
			&device.DomainTrafficData.Top20_DomainTraffic_Ranking,
			&device.DomainTrafficData.Top20_DomainTraffic_Type_Ranking,
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

func GetDeviceById(db *sql.DB,id any) (*Device, error) {
	query := `SELECT 
	id, 
	productId, 
	name, 
	uuid, 
	key, 
	cpu_usage_summary,
	memory_usage_summary,
	new_connection_summary,
	active_connection_summary,
	online_verification_summary,
	online_users_summary,
	top20UserTrafficRanking,
	top20UserTrafficGroupRanking,
	top20ServiceTrafficRanking,
	top20ServiceTrafficTypeRanking,
	top20DomainTrafficRanking,
	top20DomainTrafficTypeRanking 
	FROM devices WHERE id = ?`

	var d Device

	err := db.QueryRow(query, id).Scan(
		&d.ID, 
		&d.ProductId, 
		&d.Name, 
		&d.UUID, 
		&d.Key,
		&d.HardwareData.CPU_Usage_Summary,
		&d.HardwareData.Memory_Usage_Summary,
		&d.HardwareData.New_Connection_Summary,
		&d.HardwareData.Active_Connection_Summary,
		&d.HardwareData.Online_Verification_Summary,
		&d.HardwareData.Online_Users_Summary,
		&d.UserTrafficData.Top20_UserTraffic_Ranking,
		&d.UserTrafficData.Top20_UserTraffic_Group_Ranking,
		&d.ServiceTrafficData.Top20_ServiceTraffic_Ranking,
		&d.ServiceTrafficData.Top20_ServiceTraffic_Type_Ranking,
		&d.DomainTrafficData.Top20_DomainTraffic_Ranking,
		&d.DomainTrafficData.Top20_DomainTraffic_Type_Ranking,
	)
	
	if err != nil {
		return nil, err 
	}
	return &d, nil
}

// 新增
func CreateDevice(db *sql.DB, device Device) (int64, error) {
	query := `INSERT INTO devices (
	name, 
	productId, 
	uuid, 
	key,
	cpu_usage_summary,
	memory_usage_summary,
	new_connection_summary,
	active_connection_summary,
	online_verification_summary,
	online_users_summary,
	top20UserTrafficRanking,
	top20UserTrafficGroupRanking,
	top20ServiceTrafficRanking,
	top20ServiceTrafficTypeRanking,
	top20DomainTrafficRanking,
	top20DomainTrafficTypeRanking 
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(
		query, 
		device.Name, 
		device.ProductId,
		device.UUID, 
		device.Key,
		device.HardwareData.CPU_Usage_Summary,
		device.HardwareData.Memory_Usage_Summary,
		device.HardwareData.New_Connection_Summary,
		device.HardwareData.Active_Connection_Summary,
		device.HardwareData.Online_Verification_Summary,
		device.HardwareData.Online_Users_Summary,
		device.UserTrafficData.Top20_UserTraffic_Ranking,
		device.UserTrafficData.Top20_UserTraffic_Group_Ranking,
		device.ServiceTrafficData.Top20_ServiceTraffic_Ranking,
		device.ServiceTrafficData.Top20_ServiceTraffic_Type_Ranking,
		device.DomainTrafficData.Top20_DomainTraffic_Ranking,
		device.DomainTrafficData.Top20_DomainTraffic_Type_Ranking,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// 修改
func UpdateDevice(db *sql.DB,device Device) (any, error) {
	query := `UPDATE devices SET 
	name = ?, 
	productId = ?,
	uuid = ?, 
	key = ?,
	cpu_usage_summary = ?,
	memory_usage_summary = ?,
	new_connection_summary = ?,
	active_connection_summary = ?,
	online_verification_summary = ?,
	online_users_summary = ?,
	top20UserTrafficRanking = ?,
	top20UserTrafficGroupRanking = ?,
	top20ServiceTrafficRanking = ?,
	top20ServiceTrafficTypeRanking = ?,
	top20DomainTrafficRanking = ?,
	top20DomainTrafficTypeRanking = ?  
	WHERE id = ?`

	result, err := db.Exec(
		query, 
		device.Name, 
		device.ProductId, 
		device.UUID, 
		device.Key, 
		device.HardwareData.CPU_Usage_Summary,
		device.HardwareData.Memory_Usage_Summary,
		device.HardwareData.New_Connection_Summary,
		device.HardwareData.Active_Connection_Summary,
		device.HardwareData.Online_Verification_Summary,
		device.HardwareData.Online_Users_Summary,
		device.UserTrafficData.Top20_UserTraffic_Ranking,
		device.UserTrafficData.Top20_UserTraffic_Group_Ranking,
		device.ServiceTrafficData.Top20_ServiceTraffic_Ranking,
		device.ServiceTrafficData.Top20_ServiceTraffic_Type_Ranking,
		device.DomainTrafficData.Top20_DomainTraffic_Ranking,
		device.DomainTrafficData.Top20_DomainTraffic_Type_Ranking,
		device.ID,
	)
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