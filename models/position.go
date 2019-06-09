package models

import "database/sql"

type Position struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetAllPosition(db *sql.DB) ([]Position, error) {
	rows, err := db.Query("SELECT id, name FROM position ORDER BY name")
	if err != nil {
		return []Position{}, err
	}

	var positions []Position
	for rows.Next() {
		var position Position
		err = rows.Scan(&position.Id, &position.Name)
		if err != nil {
			return []Position{}, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}
