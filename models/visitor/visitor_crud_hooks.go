package visitor

import (
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/models"
)

func crudPreGet(id int64) error {
	return nil
}

func crudPostGet(entity *Visitor) error {
	return nil
}

func crudPreList(filters []models.ListFilter) ([]models.ListFilter, error) {
	return filters, nil
}

func crudPostList(list []*Visitor) ([]*Visitor, error) {
	return list, nil
}

func crudPreDelete(id int64, tx *sql.Tx) error {
	return nil
}

func crudPostDelete(id int64, tx *sql.Tx) error {
	return nil
}

func crudPreCreate(entity *Visitor, tx *sql.Tx) error {
	return nil
}

func crudPostCreate(entity *Visitor, tx *sql.Tx) error {
	return nil
}

func crudPreUpdate(entity *Visitor, tx *sql.Tx) error {
	return nil
}

func crudPostUpdate(entity *Visitor, tx *sql.Tx) error {
	return nil
}
