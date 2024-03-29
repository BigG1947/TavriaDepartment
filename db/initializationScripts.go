package db

const createTableDepartment = `
CREATE TABLE IF NOT EXISTS department(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	address TEXT NOT NULL UNIQUE
);`

const createTablePosition = `
CREATE TABLE IF NOT EXISTS position(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);`

const createTableEmployee = `
CREATE TABLE IF NOT EXISTS employee(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	fio TEXT NOT NULL,
	gender NUMERIC NOT NULL,
	birthday NUMERIC NOT NULL, 
	photo TEXT NOT NULL,
	address TEXT NOT NULL,
	phone TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	comment TEXT,
	date_hire NUMERIC NOT NULL,
	id_department INTEGER REFERENCES department(id) ON UPDATE CASCADE ON DELETE RESTRICT,
	id_position INTEGER REFERENCES position(id) ON UPDATE CASCADE ON DELETE RESTRICT 
);`
