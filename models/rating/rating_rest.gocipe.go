// generated by gocipe 8711bda2250e6c55fba92576b5a3047c42d034e6742f4bff5ae1f9e546907463; DO NOT EDIT

package rating

import (
	"github.com/gorilla/mux"
)

type responseSingle struct {
	Status   bool      `json:"status"`
	Messages []message `json:"messages"`
	Entity   *Rating   `json:"entity"`
}

type responseList struct {
	Status   bool      `json:"status"`
	Messages []message `json:"messages"`
	Entities []*Rating `json:"entities"`
}

type message struct {
	Type    rune   `json:"type"`
	Message string `json:"message"`
}

//RegisterRoutes registers routes with a mux Router
func RegisterRoutes(router *mux.Router) {

}