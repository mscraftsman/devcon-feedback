package feedback

import (
	"database/sql"
	"net/http"

	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/util"
)

func areResponsesInvalid(entity *Feedback) bool {
	if entity.Reaction1 == nil || entity.Reaction2 == nil || entity.Reaction3 == nil || entity.Reaction4 == nil {
		return true
	}

	switch *entity.Reaction1 {
	case "-2", "-1", "1", "2", "3":
	default:
		return true
	}

	switch *entity.Reaction2 {
	case "-2", "-1", "1", "2", "3":
	default:
		return true
	}

	switch *entity.Reaction3 {
	case "no", "yes":
	default:
		return true
	}

	return false
}

func restPreCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	visitor, err := meetup.DecodeToken(r)

	if err != nil {
		util.JSONOutputError(w, http.StatusForbidden, err.Error())
		return true, nil
	} else if entity.SessionID == nil || !sessionize.IsValidSession(*entity.SessionID) {
		util.JSONOutputError(w, http.StatusForbidden, "invalid session")
		return true, nil
	} else if areResponsesInvalid(entity) {
		util.JSONOutputError(w, http.StatusBadRequest, "incomplete feedback")
		return true, nil
	}

	*entity.VisitorID = *visitor.ID
	return false, nil
}

func restPostCreate(w http.ResponseWriter, r *http.Request, entity *Feedback, tx *sql.Tx) (bool, error) {
	//todo execute summary
	return false, nil
}
