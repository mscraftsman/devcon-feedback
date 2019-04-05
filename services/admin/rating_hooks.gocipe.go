package admin

import (
	"context"
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/auth"
	"github.com/mscraftsman/devcon-feedback/models"
)

// ratingPreRead is a hook that occurs before the ead operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPreRead(ctx context.Context, req *GetRequest, res *GetRatingResponse) (bool, error) {
	return false, nil
}

// ratingPostRead is a hook that occurs after the Read operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPostRead(ctx context.Context, req *GetRequest, res *GetRatingResponse, err error) (bool, error) {
	return false, nil
}

// ratingPreList is a hook that occurs before the ist operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPreList(ctx context.Context, req *ListRequest, res *ListRatingsResponse) (bool, error) {
	return false, nil
}

// ratingPostList is a hook that occurs after the List operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPostList(ctx context.Context, req *ListRequest, res *ListRatingsResponse, err error) (bool, error) {
	return false, nil
}

// ratingPreCreate is a hook that occurs before the reate operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPreCreate(ctx context.Context, tx *sql.Tx, req *CreateRatingRequest, res *CreateRatingResponse, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// ratingPostCreate is a hook that occurs after the Create operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPostCreate(ctx context.Context, req *CreateRatingRequest, res *CreateRatingResponse, err error) (bool, error) {
	return false, nil
}

// ratingPreUpdate is a hook that occurs before the pdate operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPreUpdate(ctx context.Context, tx *sql.Tx, req *UpdateRatingRequest, res *UpdateRatingResponse, existing models.Rating, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// ratingPostUpdate is a hook that occurs after the Update operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPostUpdate(ctx context.Context, req *UpdateRatingRequest, res *UpdateRatingResponse, existing models.Rating, err error) (bool, error) {
	return false, nil
}

// ratingPreDelete is a hook that occurs before the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPreDelete(ctx context.Context, tx *sql.Tx, req *DeleteRequest, res *DeleteResponse, existing models.Rating, err error) (bool, error) {
	return false, nil
}

// ratingPostDelete is a hook that occurs after the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) ratingPostDelete(ctx context.Context, req *DeleteRequest, res *DeleteResponse, existing models.Rating, err error) (bool, error) {
	return false, nil
}
