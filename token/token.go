package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/store"
)

var (
	// ErrorInvalidToken indicates an invalid token provided in request
	ErrorInvalidToken = errors.New("token is invalid")
)

//New creates a new token given an attendee
func New(a store.Attendee) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         a.ID,
		"name":       a.Name,
		"photo_link": a.PhotoLink,
	})

	return token.SignedString([]byte(config.JWTSecret))
}

// ToAttendee returns a Profile from a request containing jwt token
func ToAttendee(t string) (*store.Attendee, error) {
	var err error

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if token == nil || err != nil {
		return nil, ErrorInvalidToken
	}

	attendee, ok := func() (*store.Attendee, bool) {
		var (
			id          string
			idOk        bool
			name        string
			nameOk      bool
			photoLink   string
			photoLinkOk bool
		)
		claims, claimok := token.Claims.(jwt.MapClaims)

		if claimok && token.Valid {
			id, idOk = claims["id"].(string)
			name, nameOk = claims["name"].(string)
			photoLink, photoLinkOk = claims["photo_link"].(string)
		}

		return &store.Attendee{
			ID: id, Name: name, PhotoLink: photoLink,
		}, idOk && nameOk && photoLinkOk
	}()

	if !ok {
		return nil, ErrorInvalidToken
	}

	attn, err := store.DB.GetAttendee(attendee.ID)
	return &attn, err
}
