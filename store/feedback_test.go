package store_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/mscraftsman/devcon-feedback/store"
	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
)

func init() {
	// TODO: Use test suit?
	config.Load("../.env")
	sessionize.LoadSessions()

	_ = store.Init() // TODO: Catch error?
}

func TestStore_AddFeedback(t *testing.T) {
	tests := []struct {
		name string
		f    store.Feedback
		want store.Feedback
		wantErr bool
	}{
		{
			"Empty Feedback",
			store.Feedback{},
			store.Feedback{},
			true,
		},
		{
			"Missing attendeeID",
			store.Feedback{
				ID:         "0000",
				SessionID:  "1111",
				AttendeeID: "",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
				CreatedAt:  time.Now(),
			},
			store.Feedback{
				ID:         "0000",
				SessionID:  "1111",
				AttendeeID: "",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
				CreatedAt:  time.Now(),
			},
			true,
		},
		{
			"Normal Feedback",
			store.Feedback{
				ID:         "0000",
				SessionID:  "1111",
				AttendeeID: "2222",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
				CreatedAt:  time.Now(),
			},
			store.Feedback{
				ID:         "0000",
				SessionID:  "1111",
				AttendeeID: "2222",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
				CreatedAt:  time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.DB.AddFeedback(tt.f);
			if err != nil && !tt.wantErr {
				t.Errorf("Attendee.AddFeedback() Error = %v, want %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.AddFeedback() Feedback = %+v, want %+v", got, tt.want)
				return
			}
		})
	}
}

func TestStore_GetFeedback(t *testing.T) {
	// TODO: implement
}

func TestStore_ListAttendeeFeedbacks(t *testing.T) {
	// TODO: implement
}

func TestStore_ListFeedbacks(t *testing.T) {
	// TODO: implement
}
