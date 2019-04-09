package store

import (
	"sort"
	"strconv"
	"sync"
)

//Leaderboard contains aggregated stats
var Leaderboard leaderboard

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
	scores := make(Scores, len(ratings))
	set := make(map[string]interface{}) //for unique values

	for i := range ratings {
		id, score = calc(ratings[i])
		if _, ok := set[id]; !ok {
			scores[i] = Score{
				ID:    id,
				Score: score,
			}
			set[id] = nil
		}
	}

	return scores
}

func leaderboardFromRatings(ratings []Rating) leaderboard {
	var (
		speakers, sessions []Score
		wg                 sync.WaitGroup
	)
	wg.Add(2)

	go func() {
		speakers = ratingsToScores(ratings, func(r Rating) (string, int) {
			return r.Speakers, reactionScore(r.Reaction1)
		})
		sort.Sort(Scores(speakers))
		wg.Done()
	}()

	go func() {
		speakers = ratingsToScores(ratings, func(r Rating) (string, int) {
			return r.ID, reactionScore(r.Reaction2)
		})
		sort.Sort(Scores(speakers))
		wg.Done()
	}()

	wg.Wait()

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
