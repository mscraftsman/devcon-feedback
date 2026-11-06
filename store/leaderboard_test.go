package store_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
)

func TestUpdateRatingsAndLeaderboard(t *testing.T) {
	config.BoltDBPath = filepath.Join(t.TempDir(), "leaderboard.db")
	sessionize.Sessions = map[string]sessionize.Session{
		"s1": {ID: "s1", StartsAt: "2019-04-11T10:00:00", Speakers: []string{"sp1"}},
		"s2": {ID: "s2", StartsAt: "2019-04-11T11:00:00", Speakers: []string{"sp2"}},
	}
	if err := store.Init(); err != nil {
		t.Fatalf("failed to init store: %v", err)
	}
	t.Cleanup(func() {
		if store.DB != nil && store.DB.DB != nil {
			_ = store.DB.Close()
		}
	})

	if err := store.DB.SetAttendee(store.Attendee{ID: "a1", Name: "Alice", Status: true}); err != nil {
		t.Fatalf("failed to set attendee: %v", err)
	}

	if _, err := store.DB.AddFeedback(store.Feedback{
		SessionID:  "s1",
		AttendeeID: "a1",
		Reaction1:  "5",
		Reaction2:  "4",
		Reaction3:  "yes",
	}); err != nil {
		t.Fatalf("failed to add feedback: %v", err)
	}

	ratings := store.DB.ListRatings()
	if len(ratings) != 2 {
		t.Fatalf("expected 2 ratings, got %d", len(ratings))
	}

	if len(store.Leaderboard.Speakers) == 0 || store.Leaderboard.Speakers[0].ID != "s1" || store.Leaderboard.Speakers[0].Score != 5 {
		t.Errorf("unexpected top speaker leaderboard: %+v", store.Leaderboard.Speakers)
	}
	if len(store.Leaderboard.Sessions) == 0 || store.Leaderboard.Sessions[0].ID != "s1" || store.Leaderboard.Sessions[0].Score != 4 {
		t.Errorf("unexpected top session leaderboard: %+v", store.Leaderboard.Sessions)
	}
}

func TestLeaderboardTieAtLimit(t *testing.T) {
	config.BoltDBPath = filepath.Join(t.TempDir(), "leaderboard_ties.db")
	sessions := make(map[string]sessionize.Session)
	for i := 1; i <= 12; i++ {
		id := fmt.Sprintf("s%d", i)
		sessions[id] = sessionize.Session{
			ID:       id,
			StartsAt: "2019-04-11T10:00:00",
			Speakers: []string{fmt.Sprintf("sp%d", i)},
		}
	}
	sessionize.Sessions = sessions

	if err := store.Init(); err != nil {
		t.Fatalf("failed to init store: %v", err)
	}
	t.Cleanup(func() {
		if store.DB != nil && store.DB.DB != nil {
			_ = store.DB.Close()
		}
	})

	if err := store.DB.SetAttendee(store.Attendee{ID: "a1", Name: "Alice", Status: true}); err != nil {
		t.Fatalf("failed to set attendee: %v", err)
	}

	// Give s1 score 5, s2..s11 score 3 (10 sessions tied for 2nd..11th place), s12 score 1.
	for i := 1; i <= 12; i++ {
		r1 := "3"
		if i == 1 {
			r1 = "5"
		} else if i == 12 {
			r1 = "1"
		}
		if _, err := store.DB.AddFeedback(store.Feedback{
			SessionID:  fmt.Sprintf("s%d", i),
			AttendeeID: "a1",
			Reaction1:  r1,
			Reaction2:  r1,
			Reaction3:  "yes",
		}); err != nil {
			t.Fatalf("failed to add feedback for s%d: %v", i, err)
		}
	}

	// Top 10 plus the tied 11th session (score 3) => 11 entries, excluding s12 (score 1).
	if len(store.Leaderboard.Speakers) != 11 {
		t.Errorf("expected 11 speakers due to tie at rank 10, got %d: %+v", len(store.Leaderboard.Speakers), store.Leaderboard.Speakers)
	}
}
