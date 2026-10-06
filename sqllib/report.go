package sqllib

import (
	//"fmt"
	//"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Report struct {
	ID    int `json:"id"`
	DeviceId int `json:"deviceId"`
	ReportType  int `json:"reportType"`
	ReportData string `json:"reportData"`
}

// 查詢全部
func GetReport(db *sql.DB) ([]Report, error) {
	query := `SELECT 
	id, 
	deviceId,
	reportType, 
	reportData 
	FROM reports`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	

	var reports []Report

	for rows.Next() {
		var report Report

		err := rows.Scan(
			&report.ID,
			&report.DeviceId,
			&report.ReportType,
			&report.ReportData,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func GetReportById(db *sql.DB,id any) (*Report, error) {
	query := `SELECT 
	id, 
	deviceId, 
	reportType, 
	reportData 
	FROM reports WHERE id = ?`

	var r Report

	err := db.QueryRow(query, id).Scan(
		&r.ID, 
		&r.DeviceId, 
		&r.ReportType, 
		&r.ReportData,
	)
	
	if err != nil {
		return nil, err 
	}
	return &r, nil
}

// 新增
func CreateReport(db *sql.DB, report Report) (int64, error) {
	query := `INSERT INTO reports (
	deviceId, 
	reportType, 
	reportData 
	)
	VALUES (?, ?, ?)`

	result, err := db.Exec(
		query, 
		report.DeviceId, 
		report.ReportType,
		report.ReportData, 
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// 修改
func UpdateReport(db *sql.DB,report Report) (any, error) {
	query := `UPDATE reports SET 
	deviceId = ?, 
	reportType = ?,
	reportData = ? 
	WHERE id = ?`

	result, err := db.Exec(
		query, 
		report.DeviceId, 
		report.ReportType, 
		report.ReportData, 
		report.ID,
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
func DeleteReport(db *sql.DB, id string) (any, error) {
	query := `
		DELETE FROM reports
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