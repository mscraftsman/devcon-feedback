package token_test

import (
	"path/filepath"
	"testing"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
	"github.com/mscraftsman/devcon-feedback/token"
)

func TestTokenNewAndToAttendee(t *testing.T) {
	config.JWTSecret = "test-secret-key"
	config.BoltDBPath = filepath.Join(t.TempDir(), "token.db")
	sessionize.Sessions = map[string]sessionize.Session{}

	if err := store.Init(); err != nil {
		t.Fatalf("failed to init store: %v", err)
	}
	t.Cleanup(func() {
		if store.DB != nil && store.DB.DB != nil {
			_ = store.DB.Close()
		}
	})

	expected := store.Attendee{
		ID:        "42",
		Name:      "Test User",
		PhotoLink: "https://example.com/photo.jpg",
		Status:    true,
	}
	if err := store.DB.SetAttendee(expected); err != nil {
		t.Fatalf("failed to seed attendee: %v", err)
	}

	tok, err := token.New(expected)
	if err != nil {
		t.Fatalf("token.New() error: %v", err)
	}

	got, err := token.ToAttendee(tok)
	if err != nil {
		t.Fatalf("token.ToAttendee() error: %v", err)
	}
	if got.ID != expected.ID || got.Name != expected.Name || got.PhotoLink != expected.PhotoLink {
		t.Errorf("token.ToAttendee() = %+v, want %+v", got, expected)
	}

	if _, err := token.ToAttendee("invalid.token.value"); err != token.ErrorInvalidToken {
		t.Errorf("expected ErrorInvalidToken for malformed token, got %v", err)
	}
}
