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

func crudPreSave(op string, entity *Rating, tx *sql.Tx) error {
	return nil
}

func crudPostSave(op string, entity *Rating, tx *sql.Tx) error {
	return nil
}
