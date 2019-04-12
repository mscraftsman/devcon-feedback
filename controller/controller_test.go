package controller_test

import (
	"log"
	"testing"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/controller"

	"github.com/bouk/monkey"
)

func TestIsVoteOpen(t *testing.T) {
	tests := []struct {
		name         string
		want         bool
		startsAt     string
		now          string
		realtime     string
		voteClosedAt string
	}{
		{
			name:         "[1] voteClosedAt > now > startsAt",
			want:         true,
			startsAt:     "2019-04-11T12:00:00",
			now:          "2019-04-11T13:00:00",
			realtime:     "2019-04-12T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[2] voteClosedAt > now >= startsAt",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "2019-04-11T12:00:00",
			realtime:     "2019-04-12T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[3] voteClosedAt > now+16 >= startsAt",
			want:         true,
			startsAt:     "2019-04-11T12:00:00",
			now:          "2019-04-11T12:16:00",
			realtime:     "2019-04-12T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[4] now > voteClosedAt > startsAt",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "2019-04-13T18:00:00",
			realtime:     "2019-04-12T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[5] voteClosedAt > startsAt > now",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "2019-04-11T10:00:00",
			realtime:     "2019-04-12T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[6] voteClosedAt > realtime > startsAt",
			want:         true,
			startsAt:     "2019-04-11T12:00:00",
			now:          "",
			realtime:     "2019-04-11T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[7] voteClosedAt > realtime >= startsAt",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "",
			realtime:     "2019-04-11T12:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[8] voteClosedAt > realtime+16 > startsAt",
			want:         true,
			startsAt:     "2019-04-11T12:00:00",
			now:          "",
			realtime:     "2019-04-11T12:16:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[9] realtime > voteClosedAt > startsAt",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "",
			realtime:     "2019-04-13T18:30:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
		{
			name:         "[10] voteClosedAt > startsAt > realtime",
			want:         false,
			startsAt:     "2019-04-11T12:00:00",
			now:          "",
			realtime:     "2019-04-10T15:00:00",
			voteClosedAt: "2019-04-13T17:30:00",
		},
	}

	if t, err := time.LoadLocation("Indian/Mauritius"); err == nil {
		config.Tzone = t
	} else {
		log.Fatalln("Epic Failure when loading timezone")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.now == "" {
				config.Now = nil
			} else if x, err := time.Parse(config.DatetimeLayout, tt.now); err == nil {
				config.Now = &x
			} else {
				t.Errorf("Could not parse time for config.Now: %v", err)
			}

			if x, err := time.Parse(config.DatetimeLayout, tt.voteClosedAt); err == nil {
				config.VoteClosedAt = &x
			} else {
				t.Errorf("Could not parse time for config.VoteClosedAt: %v", err)
			}

			var realtime time.Time
			if x, err := time.Parse(config.DatetimeLayout, tt.realtime); err == nil {
				realtime = x
			} else {
				t.Errorf("Could not parse time for realtime: %v", err)
			}

			patch := monkey.Patch(time.Now, func() time.Time { return realtime })
			defer patch.Unpatch()

			if got := controller.IsVoteOpen(tt.startsAt); got != tt.want {
				t.Errorf("IsVoteOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}
