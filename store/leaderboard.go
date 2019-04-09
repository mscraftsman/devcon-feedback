package store

import (
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	//Leaderboard contains aggregated stats
	Leaderboard leaderboard
)

//leaderboard represens top speakers and top sessions
type leaderboard struct {
	Speakers Scores `json:"speakers"`
	Sessions Scores `json:"sessions"`
}

//Score represents an abstracted rating, whose score has been computed arbitrarily
type Score struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

//Scores represents a sortable list of scores
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s Scores) Swap(i, j int) {
	k := s[i]
	s[i] = s[j]
	s[j] = k
}

func ratingsToScores(ratings []Rating, calc func(Rating) (id string, score int)) Scores {
	var (
		id    string
		score int
	)
	scores := make(Scores, 0)
	set := make(map[string]interface{}) //for unique values

	for i, rating := range ratings {
		id, score = calc(rating)
		if id == "" {
			continue
		}

		if _, ok := set[id]; !ok {
			scores = append(scores, Score{
				ID:    id,
				Score: score,
			})
			set[id] = nil
			log.WithField("rating", rating).WithField("score", scores[i]).Debug("ratingtoscore")
		}
	}

	return scores
}

func top(scores Scores, num int) Scores {
	sort.Sort(scores)

	if len(scores) <= num {
		return scores
	}

	limited := scores[:num]
	last := limited[num-1].Score
	for i := num - 1; i < len(scores); i++ {
		if scores[i].Score == last {
			limited = append(limited, scores[i])
		}
	}

	return limited
}

func leaderboardFromRatings(ratings []Rating) leaderboard {
	speakers := top(ratingsToScores(ratings, func(r Rating) (string, int) {
		return r.ID, reactionScore(r.Reaction1)
	}), 10)

	sessions := top(ratingsToScores(ratings, func(r Rating) (string, int) {
		return r.ID, reactionScore(r.Reaction2)
	}), 10)

	return leaderboard{Speakers: speakers, Sessions: sessions}
}

func reactionScore(reactions map[string]ReactionSummary) int {
	var v, t int

	for i := range reactions {
		v, _ = strconv.Atoi(reactions[i].Reaction)
		t = v * reactions[i].Count
	}

	return t
}
