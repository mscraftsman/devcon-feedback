package store

import (
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//Leaderboard contains aggregated stats
var Leaderboard leaderboard

//leaderboard represens top speakers and top sessions
type leaderboard struct {
	Speakers []string
	Sessions []string
}

func leaderboardFromRatings(ratings []Rating) leaderboard {
	var (
		l leaderboard

		spk = make(chan []string)
		ses = make(chan []string)
	)

	log.Debug("leaderboardFromRatings:start")

	go topSpeakerRatings(ratings, 3, spk)
	go topSessionRatings(ratings, 3, ses)

	l.Speakers = <-spk
	l.Sessions = <-ses

	close(spk)
	close(ses)

	log.WithField("leaderboard", l).Debug("leaderboardFromRatings:done")

	return l
}

func reactionScore(reactions map[string]ReactionSummary) int {
	var v, t int

	for i := range reactions {
		v, _ = strconv.Atoi(reactions[i].Reaction)
		t = v * reactions[i].Count
	}

	return t
}

type ratingsBySpeaker []Rating

func (l ratingsBySpeaker) Len() int {
	return len(l)
}

func (l ratingsBySpeaker) Less(i, j int) bool {
	return reactionScore(l[i].Reaction1) < reactionScore(l[j].Reaction1)
}

func (l ratingsBySpeaker) Swap(i, j int) {
	k := l[i]
	l[i] = l[j]
	l[j] = k
}

func topSpeakerRatings(ratings []Rating, count int, done chan<- []string) {
	var top []Rating
	sort.Sort(ratingsBySpeaker(ratings))

	if len(ratings) == count {
		top = ratings
		log.WithField("count", count).Debug("topSpeakerRatings")
	} else {
		top = ratings[:count]

		//ex-aequo
		lastScore := reactionScore(ratings[count-1].Reaction1)
		for i := count - 1; i < len(ratings); i++ {
			if reactionScore(ratings[i].Reaction1) == lastScore {
				top = append(top, ratings[i])
			} else {
				break
			}
		}

		log.WithField("count", len(top)).Debug("topSpeakerRatings")
	}

	var out []string
	for i := range top {
		out = append(out, top[i].Speakers)
	}

	done <- out
}

type ratingsBySession []Rating

func (l ratingsBySession) Len() int {
	return len(l)
}

func (l ratingsBySession) Less(i, j int) bool {
	return reactionScore(l[i].Reaction2) < reactionScore(l[j].Reaction2)
}

func (l ratingsBySession) Swap(i, j int) {
	k := l[i]
	l[i] = l[j]
	l[j] = k
}

func topSessionRatings(ratings []Rating, count int, done chan<- []string) {
	var top []Rating
	sort.Sort(ratingsBySession(ratings))

	if len(ratings) == count {
		top = ratings
		log.WithField("count", count).Debug("topSessionRatings")
	} else {
		top = ratings[:count]

		//ex-aequo
		lastScore := reactionScore(ratings[count-1].Reaction2)
		for i := count - 1; i < len(ratings); i++ {
			if reactionScore(ratings[i].Reaction2) == lastScore {
				top = append(top, ratings[i])
			} else {
				break
			}
		}

		log.WithField("count", len(top)).Debug("topSpeakerRatings")
	}

	var out []string
	for i := range top {
		out = append(out, top[i].ID)
	}

	done <- out
}
