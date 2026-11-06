package store_test

import (
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
)

func setupTestDB(t *testing.T) {
	t.Helper()
	config.BoltDBPath = filepath.Join(t.TempDir(), "test.db")
	sessionize.Sessions = map[string]sessionize.Session{
		"1111": {
			ID:       "1111",
			StartsAt: "2019-04-11T10:00:00",
			Speakers: []string{"speaker-1"},
		},
	}
	if err := store.Init(); err != nil {
		t.Fatalf("failed to initialize test store: %v", err)
	}
	t.Cleanup(func() {
		if store.DB != nil && store.DB.DB != nil {
			_ = store.DB.Close()
		}
	})
}

func TestStore_AddFeedback(t *testing.T) {
	setupTestDB(t)

	if err := store.DB.SetAttendee(store.Attendee{
		ID:     "2222",
		Name:   "Jane Doe",
		Status: true,
	}); err != nil {
		t.Fatalf("failed to seed attendee: %v", err)
	}

	now := time.Now()
	tests := []struct {
		name    string
		f       store.Feedback
		want    store.Feedback
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
				CreatedAt:  now,
			},
			store.Feedback{
				ID:         "0000",
				SessionID:  "1111",
				AttendeeID: "",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
				CreatedAt:  now,
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
				CreatedAt:  now,
			},
			store.Feedback{
				ID:         "1111:2222",
				SessionID:  "1111",
				AttendeeID: "2222",
				Reaction1:  "3",
				Reaction2:  "2",
				Reaction3:  "yes",
				Reaction4:  "blablabla",
			},
			false,
		},
		{
			"Duplicate Feedback",
			store.Feedback{
				SessionID:  "1111",
				AttendeeID: "2222",
				Reaction1:  "4",
				Reaction2:  "5",
				Reaction3:  "yes",
				Reaction4:  "duplicate",
			},
			store.Feedback{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.DB.AddFeedback(tt.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.AddFeedback() Error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				tt.want.CreatedAt = got.CreatedAt
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Store.AddFeedback() Feedback = %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestStore_GetFeedback(t *testing.T) {
	setupTestDB(t)

	if err := store.DB.SetAttendee(store.Attendee{
		ID:     "2222",
		Name:   "Jane Doe",
		Status: true,
	}); err != nil {
		t.Fatalf("failed to seed attendee: %v", err)
	}

	added, err := store.DB.AddFeedback(store.Feedback{
		SessionID:  "1111",
		AttendeeID: "2222",
		Reaction1:  "5",
		Reaction2:  "4",
		Reaction3:  "yes",
		Reaction4:  "great session",
	})
	if err != nil {
		t.Fatalf("unexpected error adding feedback: %v", err)
	}

	got, err := store.DB.GetFeedback(added.ID)
	if err != nil {
		t.Fatalf("Store.GetFeedback() unexpected error: %v", err)
	}
	if got.ID != added.ID || got.Reaction4 != "great session" {
		t.Errorf("Store.GetFeedback() = %+v, want %+v", got, added)
	}

	if _, err := store.DB.GetFeedback("nonexistent"); err != store.ErrorFeedbackNotFound {
		t.Errorf("Store.GetFeedback(nonexistent) err = %v, want %v", err, store.ErrorFeedbackNotFound)
	}
}

func TestStore_ListAttendeeFeedbacks(t *testing.T) {
	setupTestDB(t)

	if err := store.DB.SetAttendee(store.Attendee{
		ID:     "2222",
		Name:   "Jane Doe",
		Status: true,
	}); err != nil {
		t.Fatalf("failed to seed attendee: %v", err)
	}

	if _, err := store.DB.AddFeedback(store.Feedback{
		SessionID:  "1111",
		AttendeeID: "2222",
		Reaction1:  "5",
		Reaction2:  "4",
		Reaction3:  "yes",
		Reaction4:  "awesome",
	}); err != nil {
		t.Fatalf("unexpected error adding feedback: %v", err)
	}

	list, err := store.DB.ListAttendeeFeedbacks("2222")
	if err != nil {
		t.Fatalf("Store.ListAttendeeFeedbacks() unexpected error: %v", err)
	}
	if len(list) != 1 || list[0].ID != "1111:2222" {
		t.Errorf("Store.ListAttendeeFeedbacks() = %+v, want 1 item with ID 1111:2222", list)
	}

	if _, err := store.DB.ListAttendeeFeedbacks("unknown"); err != store.ErrorAttendeeNotFound {
		t.Errorf("Store.ListAttendeeFeedbacks(unknown) err = %v, want %v", err, store.ErrorAttendeeNotFound)
	}
}

func TestStore_ListFeedbacks(t *testing.T) {
	setupTestDB(t)

	if err := store.DB.SetAttendee(store.Attendee{
		ID:     "2222",
		Name:   "Jane Doe",
		Status: true,
	}); err != nil {
		t.Fatalf("failed to seed attendee: %v", err)
	}

	if _, err := store.DB.AddFeedback(store.Feedback{
		SessionID:  "1111",
		AttendeeID: "2222",
		Reaction1:  "4",
		Reaction2:  "4",
		Reaction3:  "no",
		Reaction4:  "good",
	}); err != nil {
		t.Fatalf("unexpected error adding feedback: %v", err)
	}

	all := store.DB.ListFeedbacks()
	if len(all) != 1 || all[0].ID != "1111:2222" {
		t.Errorf("Store.ListFeedbacks() = %+v, want 1 item with ID 1111:2222", all)
	}
}
