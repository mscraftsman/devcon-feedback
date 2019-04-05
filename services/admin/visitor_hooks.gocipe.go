package admin

import (
	"context"
	"database/sql"

	"github.com/mscraftsman/devcon-feedback/auth"
	"github.com/mscraftsman/devcon-feedback/models"
)

// visitorPreRead is a hook that occurs before the ead operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPreRead(ctx context.Context, req *GetRequest, res *GetVisitorResponse) (bool, error) {
	return false, nil
}

// visitorPostRead is a hook that occurs after the Read operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPostRead(ctx context.Context, req *GetRequest, res *GetVisitorResponse, err error) (bool, error) {
	return false, nil
}

// visitorPreList is a hook that occurs before the ist operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPreList(ctx context.Context, req *ListRequest, res *ListVisitorsResponse) (bool, error) {
	return false, nil
}

// visitorPostList is a hook that occurs after the List operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPostList(ctx context.Context, req *ListRequest, res *ListVisitorsResponse, err error) (bool, error) {
	return false, nil
}

// visitorPreCreate is a hook that occurs before the reate operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPreCreate(ctx context.Context, tx *sql.Tx, req *CreateVisitorRequest, res *CreateVisitorResponse, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// visitorPostCreate is a hook that occurs after the Create operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPostCreate(ctx context.Context, req *CreateVisitorRequest, res *CreateVisitorResponse, err error) (bool, error) {
	return false, nil
}

// visitorPreUpdate is a hook that occurs before the pdate operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPreUpdate(ctx context.Context, tx *sql.Tx, req *UpdateVisitorRequest, res *UpdateVisitorResponse, existing models.Visitor, passport auth.Passport, err error) (bool, error) {
	return false, nil
}

// visitorPostUpdate is a hook that occurs after the Update operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPostUpdate(ctx context.Context, req *UpdateVisitorRequest, res *UpdateVisitorResponse, existing models.Visitor, err error) (bool, error) {
	return false, nil
}

// visitorPreDelete is a hook that occurs before the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPreDelete(ctx context.Context, tx *sql.Tx, req *DeleteRequest, res *DeleteResponse, existing models.Visitor, err error) (bool, error) {
	return false, nil
}

// visitorPostDelete is a hook that occurs after the Delete operation in Admin service. Returns stop (boolean) and error
func (s Service) visitorPostDelete(ctx context.Context, req *DeleteRequest, res *DeleteResponse, existing models.Visitor, err error) (bool, error) {
	return false, nil
}
