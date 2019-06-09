package models

import (
	"database/sql"
	"fmt"
)

const (
	FioField          = "fio"
	GenderField       = "gender"
	BirthdayField     = "birthday"
	PhotoField        = "photo"
	AddressField      = "address"
	PhoneField        = "phone"
	EmailField        = "email"
	CommentField      = "comment"
	DateHireField     = "date_hire"
	IdDepartmentField = "id_department"
	IdPositionField   = "id_position"
)

type Employee struct {
	Id         int64      `json:"id"`
	Fio        string     `json:"fio"`
	Gender     bool       `json:"gender"`
	Birthday   string     `json:"birthday"`
	Photo      string     `json:"photo"`
	Address    string     `json:"address"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	Comment    string     `json:"comment"`
	DateHire   string     `json:"date_hire"`
	Department Department `json:"department"`
	Position   Position   `json:"position"`
}

func GetAllEmployeeFromDepartment(db *sql.DB, idDepartment int64) ([]Employee, error) {
	rows, err := db.Query("SELECT e.id, e.fio, e.gender, e.birthday, e.photo, e.address, e.phone, e.email, e.comment, e.date_hire, d.id, d.address, p.id, p.name FROM employee e INNER JOIN department d on e.id_department = d.id INNER JOIN position p on e.id_position = p.id WHERE e.id_department = ? ORDER BY e.fio", idDepartment)
	if err != nil {
		return []Employee{}, err
	}
	var employees []Employee

	for rows.Next() {
		var employee Employee
		var comment sql.NullString
		err = rows.Scan(&employee.Id, &employee.Fio, &employee.Gender, &employee.Birthday, &employee.Photo, &employee.Address, &employee.Phone, &employee.Email, &comment, &employee.DateHire, &employee.Department.Id, &employee.Department.Address, &employee.Position.Id, &employee.Position.Name)
		if err != nil {
			return []Employee{}, err
		}
		if comment.Valid {
			employee.Comment = comment.String
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func GetEmployeeById(db *sql.DB, id int64) (*Employee, error) {
	var employee Employee
	var comment sql.NullString

	err := db.QueryRow("SELECT e.id, e.fio, e.gender, e.birthday, e.photo, e.address, e.phone, e.email, e.comment, e.date_hire, d.id, d.address, p.id, p.name FROM employee e INNER JOIN department d on e.id_department = d.id INNER JOIN position p on e.id_position = p.id WHERE e.id = ?", id).Scan(&employee.Id, &employee.Fio, &employee.Gender, &employee.Birthday, &employee.Photo, &employee.Address, &employee.Phone, &employee.Email, &comment, &employee.DateHire, &employee.Department.Id, &employee.Department.Address, &employee.Position.Id, &employee.Position.Name)

	if err != nil {
		return &Employee{}, err
	}
	if comment.Valid {
		employee.Comment = comment.String
	}

	return &employee, nil
}

func AddEmployee(db *sql.DB, employee *Employee) (int64, error) {
	res, err := db.Exec("INSERT INTO employee(fio, gender, birthday, photo, address, phone, email, comment, date_hire, id_department, id_position) VALUES (?,?,?,?,?,?,?,?,?,?,?)", employee.Fio, employee.Gender, employee.Birthday, employee.Photo, employee.Address, employee.Phone, employee.Email, employee.Comment, employee.DateHire, employee.Department.Id, employee.Position.Id)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func DeleteEmployee(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM employee WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(db *sql.DB, idEmployee int64, params map[string]string) error {
	lenData := len(params)
	if lenData == 0 {
		return nil
	}
	sqlScript := fmt.Sprintf("UPDATE employee SET")
	var i int
	for key, value := range params {
		if i < lenData-1 {
			sqlScript += fmt.Sprintf(" %s = '%s',", key, value)
		} else if i == lenData-1 {
			sqlScript += fmt.Sprintf(" %s = '%s'", key, value)
		}
		i++
	}
	sqlScript += fmt.Sprintf(" WHERE employee.id = ?;")
	_, err := db.Exec(sqlScript, idEmployee)
	if err != nil {
		return err
	}
	return nil
}
