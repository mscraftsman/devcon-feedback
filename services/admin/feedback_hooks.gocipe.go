package admin

import (
	"context"
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/auth"
	"github.com/mscraftsman/devcon-feedback/models"
)

// feedbackPreRead is a hook that occurs before the ead operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPreRead(ctx context.Context, req *GetRequest, res *GetFeedbackResponse) (bool, error) {
	return false, nil
}

// feedbackPostRead is a hook that occurs after the Read operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPostRead(ctx context.Context, req *GetRequest, res *GetFeedbackResponse, err error) (bool, error) {
	return false, nil
}

// feedbackPreList is a hook that occurs before the ist operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPreList(ctx context.Context, req *ListRequest, res *ListFeedbacksResponse) (bool, error) {
	return false, nil
}

// feedbackPostList is a hook that occurs after the List operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPostList(ctx context.Context, req *ListRequest, res *ListFeedbacksResponse, err error) (bool, error) {
	return false, nil
}

// feedbackPreCreate is a hook that occurs before the reate operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPreCreate(ctx context.Context, tx *sql.Tx, req *CreateFeedbackRequest, res *CreateFeedbackResponse, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// feedbackPostCreate is a hook that occurs after the Create operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPostCreate(ctx context.Context, req *CreateFeedbackRequest, res *CreateFeedbackResponse, err error) (bool, error) {
	return false, nil
}

// feedbackPreUpdate is a hook that occurs before the pdate operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPreUpdate(ctx context.Context, tx *sql.Tx, req *UpdateFeedbackRequest, res *UpdateFeedbackResponse, existing models.Feedback, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// feedbackPostUpdate is a hook that occurs after the Update operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPostUpdate(ctx context.Context, req *UpdateFeedbackRequest, res *UpdateFeedbackResponse, existing models.Feedback, err error) (bool, error) {
	return false, nil
}

// feedbackPreDelete is a hook that occurs before the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPreDelete(ctx context.Context, tx *sql.Tx, req *DeleteRequest, res *DeleteResponse, existing models.Feedback, err error) (bool, error) {
	return false, nil
}

// feedbackPostDelete is a hook that occurs after the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) feedbackPostDelete(ctx context.Context, req *DeleteRequest, res *DeleteResponse, existing models.Feedback, err error) (bool, error) {
	return false, nil
}
