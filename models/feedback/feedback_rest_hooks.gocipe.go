package feedback

import (
	"database/sql"
	"net/http"
)

func restPreCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}

func restPostCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	return false, nil
}
