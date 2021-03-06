package store

import (
	"encoding/json"
	"strconv"

	"github.com/mscraftsman/devcon-feedback/sessionize"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

//Rating for each session; id is same as sessionid
type Rating struct {
	ID        string                     `json:"id"`
	Score     int                        `json:"score"`
	Count     int                        `json:"count"`
	Reaction1 map[string]ReactionSummary `json:"reaction1"`
	Reaction2 map[string]ReactionSummary `json:"reaction2"`
	Reaction3 map[string]ReactionSummary `json:"reaction3"`
}

//ReactionSummary represents an aggregate of all reactions
type ReactionSummary struct {
	Reaction string `json:"reaction"`
	Count    int    `json:"count"`
}

//UpdateRatings recalculate all ratings
func (s *Store) UpdateRatings() {
	defer func() {
		if r := recover(); r != nil {
			log.WithField("recovered", r).Error("update:ratings:panic")
		}
	}()

	mRatings := make(map[string]Rating)
	for _, session := range sessionize.Sessions {
		if len(session.Speakers) != 0 {
			mRatings[session.ID] = Rating{
				ID:        session.ID,
				Reaction1: make(map[string]ReactionSummary),
				Reaction2: make(map[string]ReactionSummary),
				Reaction3: make(map[string]ReactionSummary),
			}
		}
	}

	isBanned := s.IsAttendeeBanned()
	for _, feedback := range s.ListFeedbacks() {
		if !isBanned(feedback.AttendeeID) {
			mRatings[feedback.SessionID] = computeRating(mRatings[feedback.SessionID], feedback)
		}
	}

	var ratings []Rating
	_ = s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketRatings)

		for _, mRating := range mRatings {
			if mRatingJSON, err := json.Marshal(mRating); err == nil {
				if err := b.Put([]byte(mRating.ID), mRatingJSON); err != nil {
					log.WithField("error", err).Error("ratings:update:save")
				} else {
					ratings = append(ratings, mRating)
				}
			}
		}

		return nil
	})

	Leaderboard = leaderboardFromRatings(ratings)
}

//ListRatings returns all ratings
func (s *Store) ListRatings() []Rating {
	var (
		ratings []Rating
		err     error
	)

	_ = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketRatings)
		_ = b.ForEach(func(k, v []byte) error {
			var r Rating
			if err = json.Unmarshal(v, &r); err == nil {
				ratings = append(ratings, r)
				log.WithField("rating", r).Debug("ratings:list")
			}
			return nil
		})

		return nil
	})

	return ratings
}

func computeRating(rating Rating, feedback Feedback) Rating {
	r1 := feedback.Reaction1
	v1, _ := strconv.Atoi(r1)
	if reaction1, ok := rating.Reaction1[r1]; ok {
		reaction1.Count++
		rating.Reaction1[r1] = reaction1
	} else {
		rating.Reaction1[r1] = ReactionSummary{
			Reaction: r1,
			Count:    1,
		}
	}

	r2 := feedback.Reaction2
	v2, _ := strconv.Atoi(r2)
	if reaction2, ok := rating.Reaction2[r2]; ok {
		reaction2.Count++
		rating.Reaction2[r2] = reaction2
	} else {
		rating.Reaction2[r2] = ReactionSummary{
			Reaction: r2,
			Count:    1,
		}
	}

	rating.Count++
	rating.Score += v1 + v2
	log.WithField("rating", rating).Debug("ratings:compute")
	return rating
}
