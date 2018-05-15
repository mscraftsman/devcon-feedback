package feedback

import (
	"database/sql"
	"net/http"

	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/util"
)

func restPreCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	visitor, err := meetup.DecodeToken(r)

	if err != nil {
		util.JSONOutputError(w, http.StatusForbidden, err.Error())
		return true, nil
	} else if entity.SessionID == nil || !sessionize.IsValidSession(*entity.SessionID) {
		util.JSONOutputError(w, http.StatusForbidden, "invalid session")
		return true, nil
	}

	*entity.VisitorID = *visitor.ID
	return false, nil
}

func restPostCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	//todo execute summary
	return false, nil
}
