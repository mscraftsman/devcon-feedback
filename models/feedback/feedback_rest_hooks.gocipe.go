package feedback

import (
	"database/sql"
	"net/http"

	"github.com/mscraftsman/devcon-feedback/models"
)

func restPreGet(w http.ResponseWriter, r *http.Request, id int64) (bool, error) {
	return false, nil
}

func restPostGet(w http.ResponseWriter, r *http.Request, entity *Feedback) (bool, error) {
	return false, nil
}

func restPreList(w http.ResponseWriter, r *http.Request, filters []models.ListFilter) ([]models.ListFilter, bool, error) {
	return filters, false, nil
}

func restPostList(w http.ResponseWriter, r *http.Request, list []*Feedback) ([]*Feedback, bool, error) {
	return list, false, nil
}

func restPreCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPostCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPreUpdate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPostUpdate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPreDelete(w http.ResponseWriter, r *http.Request, id int64, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPostDelete(w http.ResponseWriter, r *http.Request, id int64, tx *sql.Tx) (bool, error) {
	return false, nil
}
