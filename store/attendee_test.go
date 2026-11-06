package store_test

import (
	"reflect"
	"testing"

	"github.com/mscraftsman/devcon-feedback/store"
)

func TestAttendee_AddBookmark(t *testing.T) {
	tests := []struct {
		name string
		a    store.Attendee
		id   string
		want store.Attendee
	}{
		{
			"Empty attendee.Bookmarks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
		},
		{
			"Duplicate Bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
		},
		{
			"Non-empty attendee.Bookmarks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			"2345678",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
			},
		},
		{
			"Empty Bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			"",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
		},
		// {
		// 	"Invalid Bookmark ID (Session doesn't exist)",
		// 	store.Attendee{
		// 		ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
		// 	},
		// 	"1111111",
		// 	store.Attendee{
		// 		ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddBookmark(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.AddBookmark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendee_RemoveBookmark(t *testing.T) {
	tests := []struct {
		name string
		a    store.Attendee
		id   string
		want store.Attendee
	}{
		{
			"Empty attendee.Bookmarks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
		},
		{
			"From 1 bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
		},
		{
			"From 2 bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "2345678", Feedbacks: "",
			},
		},
		{
			"Empty Bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			"",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
		},
		{
			"Bookmark not in attendee.Bookmarks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
			},
			"1111111",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.RemoveBookmark(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.RemoveBookmark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendee_ListBookmarks(t *testing.T) {
	tests := []struct {
		name string
		a    store.Attendee
		want []string
	}{
		{
			"Empty attendee.Bookmarks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
			[]string{},
		},
		{
			"From 1 bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			[]string{"1234567"},
		},
		{
			"From 2 bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567;2345678", Feedbacks: "",
			},
			[]string{"1234567", "2345678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ListBookmarks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.ListBookmarks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendee_AddFeedback(t *testing.T) {
	tests := []struct {
		name string
		a    store.Attendee
		id   string
		want store.Attendee
	}{
		{
			"Empty attendee.Feedbacks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
		},
		{
			"Duplicate Feedback",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
			"1234567",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
		},
		{
			"Non-empty attendee.Feedbacks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
			"2345678",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567;2345678",
			},
		},
		{
			"Empty Bookmark",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
			"",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
		},
		// {
		// 	"Invalid Feedback ID (Session doesn't exist)",
		// 	store.Attendee{
		// 		ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "",
		// 	},
		// 	"1111111",
		// 	store.Attendee{
		// 		ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567;2345678",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AddFeedback(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.AddFeedback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendee_ListFeedbacks(t *testing.T) {
	tests := []struct {
		name string
		a    store.Attendee
		want []string
	}{
		{
			"Empty attendee.Feedbacks",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "1234567", Feedbacks: "",
			},
			[]string{},
		},
		{
			"From 1 feedback",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567",
			},
			[]string{"1234567"},
		},
		{
			"From 2 feedback",
			store.Attendee{
				ID: "000", Name: "John Doe", PhotoLink: "http://foo.bar", Status: true, Bookmarks: "", Feedbacks: "1234567;2345678",
			},
			[]string{"1234567", "2345678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ListFeedbacks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attendee.ListFeedbacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
