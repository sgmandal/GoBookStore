package helper

import (
	"bookstore/database"
	"bookstore/models"
)

func GetEveryBook() ([]models.Book, error) {
	xd := []models.Book{}
	err := database.Db.Select(&xd, `SELECT * FROM Books`)
	if err != nil {
		return nil, err
	}

	return xd, nil
}
