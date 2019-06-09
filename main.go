package main

import (
	"fmt"
	"log"
	"math/rand"
	db2 "tavriaDepartment/db"
	"tavriaDepartment/models"
	"time"
)

var db *db2.DB

func main() {
	var err error
	db, err = db2.Init()
	if err != nil {
		log.Printf("Error in initialization DB: %s\n", err)
		return
	}
	tests()
	return
}

func tests() {

	id, err := models.AddDepartment(db.Connection, GetRandomString())
	if err != nil {
		log.Printf("Department add error:\n%s\n", err)
		return
	}
	fmt.Printf("Department add successful\n")

	department, err := models.GetDepartmentById(db.Connection, id)
	if err != nil {
		log.Printf("GetDepartmentById error:\n%s\n", err)
		return
	}
	fmt.Printf("Department get by id successful\n%v\n", department)

	err = models.UpdateDepartment(db.Connection, department.Id, GetRandomString())
	if err != nil {
		log.Printf("UpdateDepartment error:\n%s\n", err)
		return
	}
	department, _ = models.GetDepartmentById(db.Connection, department.Id)
	fmt.Printf("Department update successful\n%v\n", department)

	err = models.DeleteDepartment(db.Connection, department.Id)
	if err != nil {
		log.Printf("Department delete error:\n%s\n", err)
		return
	}
	fmt.Printf("Department delete successful\n")

	var employee1 models.Employee
	employee1.Department.Id = 1
	employee1.Address = GetRandomString()
	employee1.Position.Id = 1
	employee1.DateHire = "2019-02-03"
	employee1.Phone = GetRandomString()
	employee1.Birthday = "2019-02-02"
	employee1.Gender = true
	employee1.Fio = GetRandomString()
	employee1.Email = GetRandomString()
	employee1.Photo = GetRandomString()

	id, err = models.AddEmployee(db.Connection, &employee1)
	if err != nil {
		log.Printf("AddEmployee error:\n%s\n", err)
		return
	}
	fmt.Printf("AddEmployee successful\n")

	employee2, err := models.GetEmployeeById(db.Connection, id)
	if err != nil {
		log.Printf("GetEmployeeById error:\n%s\n", err)
		return
	}
	fmt.Printf("GetEmployeeById successful\n%#v\n", employee2)

	params := make(map[string]string)
	params["fio"] = "UpdateUser_" + GetRandomString()
	params["gender"] = "false"
	err = models.UpdateEmployee(db.Connection, employee2.Id, params)
	if err != nil {
		log.Printf("Update Employee error:\n%s\n", err)
		return
	}
	employee2, err = models.GetEmployeeById(db.Connection, employee2.Id)
	fmt.Printf("Update Employee successful\n%#v\n", employee2)

	err = models.DeleteEmployee(db.Connection, employee2.Id)
	if err != nil {
		log.Printf("Delete Employee error:\n%s\n", err)
		return
	}

	fmt.Printf("Delete Employee sucessfull\n")

	employees, err := models.GetAllEmployeeFromDepartment(db.Connection, 1)
	if err != nil {
		log.Printf("GetAllEmployeeFromDepartment error:\n%s\n", err)
		return
	}
	fmt.Printf("\n\nGetAllEmployeeWithDepartment\n\n")

	for _, e := range employees {
		fmt.Printf("%#v\n", e)
	}
	return
}

func GetRandomString() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
