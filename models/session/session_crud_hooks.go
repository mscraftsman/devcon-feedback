package session

import (
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/models"
)

func crudPreGet(id string) error {
	return nil
}

func crudPostGet(entity *Session) error {
	return nil
}

func crudPreList(filters []models.ListFilter) ([]models.ListFilter, error) {
	return filters, nil
}

func crudPostList(list []*Session) ([]*Session, error) {
	return list, nil
}

func crudPreDelete(id string, tx *sql.Tx) error {
	return nil
}

func crudPostDelete(id string, tx *sql.Tx) error {
	return nil
}

func crudPreCreate(entity *Session, tx *sql.Tx) error {
	return nil
}

func crudPostCreate(entity *Session, tx *sql.Tx) error {
	return nil
}

func crudPreUpdate(entity *Session, tx *sql.Tx) error {
	return nil
}

func crudPostUpdate(entity *Session, tx *sql.Tx) error {
	return nil
}