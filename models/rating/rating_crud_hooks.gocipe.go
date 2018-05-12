package rating

import (
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/models"
)

func crudPreGet(id string) error {
	return nil
}

func crudPostGet(entity *Rating) error {
	return nil
}

func crudPreList(filters []models.ListFilter) ([]models.ListFilter, error) {
	return filters, nil
}

func crudPostList(list []*Rating) ([]*Rating, error) {
	return list, nil
}

func crudPreDelete(id string, tx *sql.Tx) error {
	return nil
}

func crudPostDelete(id string, tx *sql.Tx) error {
	return nil
}

func crudPreCreate(entity *Rating, tx *sql.Tx) error {
	return nil
}

func crudPostCreate(entity *Rating, tx *sql.Tx) error {
	return nil
}

func crudPreUpdate(entity *Rating, tx *sql.Tx) error {
	return nil
}

func crudPostUpdate(entity *Rating, tx *sql.Tx) error {
	return nil
}
