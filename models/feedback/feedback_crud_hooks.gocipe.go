package feedback

import (
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/models"
)

func crudPreGet(id int64) error {
	return nil
}

func crudPostGet(entity *Feedback) error {
	return nil
}

func crudPreList(filters []models.ListFilter) ([]models.ListFilter, error) {
	return filters, nil
}

func crudPostList(list []*Feedback) ([]*Feedback, error) {
	return list, nil
}

func crudPreDelete(id int64, tx *sql.Tx) error {
	return nil
}

func crudPostDelete(id int64, tx *sql.Tx) error {
	return nil
}

func crudPreSave(op string, entity *Feedback, tx *sql.Tx) error {
	return nil
}

func crudPostSave(op string, entity *Feedback, tx *sql.Tx) error {
	return nil
}
