package models

import "database/sql"

type Department struct {
	Id      int64  `json:"id"`
	Address string `json:"address"`
}

func GetAllDepartments(db *sql.DB) ([]Department, error) {
	rows, err := db.Query("SELECT id, address FROM department ORDER BY address ASC")
	if err != nil {
		return []Department{}, err
	}

	var departments []Department
	for rows.Next() {
		var department Department
		err = rows.Scan(&department.Id, &department.Address)
		if err != nil {
			return []Department{}, err
		}
		departments = append(departments, department)
	}
	return departments, nil
}

func GetDepartmentById(db *sql.DB, id int64) (*Department, error) {
	var department Department
	err := db.QueryRow("SELECT id, address FROM department WHERE id = ?", id).Scan(&department.Id, &department.Address)
	if err != nil {
		return &Department{}, err
	}
	return &department, err
}

func AddDepartment(db *sql.DB, address string) (int64, error) {
	res, err := db.Exec("INSERT INTO department(address) VALUES (?)", address)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateDepartment(db *sql.DB, idDepartment int64, address string) error {
	_, err := db.Exec("UPDATE department SET address = ? WHERE id = ?", address, idDepartment)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(db *sql.DB, idDepartment int64) error {
	_, err := db.Exec("DELETE FROM department WHERE id = ?", idDepartment)
	if err != nil {
		return err
	}
	return nil
}
