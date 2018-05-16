package feedback

import (
	"net/http"

	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/util"

	"github.com/mscraftsman/devcon-feedback/models"
)

// MyFeedback is an endpoint for /b/feedbacks/me
func MyFeedback(w http.ResponseWriter, r *http.Request) {
	v, err := meetup.DecodeToken(r)

	if err != nil {
		util.JSONOutputError(w, http.StatusForbidden, err.Error())
		return
	}

	f, err := List([]models.ListFilter{
		models.ListFilter{
			Field:     "visitor_id",
			Operation: "=",
			Value:     *v.ID,
		},
	})

	if err != nil {
		util.JSONOutputError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSONOutputResponse(w, f)
}
