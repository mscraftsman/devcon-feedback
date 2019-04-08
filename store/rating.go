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

	ratings := make(map[string]Rating)
	banned := func() map[string]interface{} {
		a := s.ListAttendees()
		b := make(map[string]interface{})

		for i := range a {
			if !a[i].Status {
				b[a[i].ID] = nil
			}
		}

		return b
	}()

	feedbacks := s.ListFeedbacks()
	log.WithField("banned", banned).WithField("feedbacks", feedbacks).Debug("ratings:update")

	for sid := range sessionize.Sessions {
		ratings[sid] = Rating{
			ID:        sid,
			Reaction1: make(map[string]ReactionSummary),
			Reaction2: make(map[string]ReactionSummary),
			Reaction3: make(map[string]ReactionSummary),
		}
	}

	for i := range feedbacks {
		if _, ok := banned[feedbacks[i].AttendeeID]; ok {
			continue
		}

		rating := ratings[feedbacks[i].SessionID]

		r1 := feedbacks[i].Reaction1
		v1, _ := strconv.Atoi(r1)
		if _, ok := rating.Reaction1[r1]; ok {
			r := rating.Reaction1[r1]
			r.Count++
			rating.Reaction1[r1] = r
		} else {
			rating.Reaction1[r1] = ReactionSummary{
				Reaction: r1,
				Count:    1,
			}
		}

		r2 := feedbacks[i].Reaction2
		v2, _ := strconv.Atoi(r2)
		if _, ok := rating.Reaction2[r2]; ok {
			r := rating.Reaction2[r2]
			r.Count++
			rating.Reaction2[r2] = r
		} else {
			rating.Reaction2[r2] = ReactionSummary{
				Reaction: r2,
				Count:    1,
			}
		}

		r3 := feedbacks[i].Reaction3
		v3, _ := strconv.Atoi(r3)
		if _, ok := rating.Reaction3[r3]; ok {
			r := rating.Reaction3[r3]
			r.Count++
			rating.Reaction3[r3] = r
		} else {
			rating.Reaction3[r3] = ReactionSummary{
				Reaction: r3,
				Count:    1,
			}
		}

		rating.Count++
		rating.Score += v1 + v2 + v3
		ratings[feedbacks[i].SessionID] = rating
		log.WithField("rating", rating).Debug("ratings:update")
	}

	s.Update(func(tx *bolt.Tx) error {
		var (
			err error
			j   []byte
		)
		b := tx.Bucket(bucketRatings)

		for i := range ratings {
			if j, err = json.Marshal(ratings[i]); err == nil {
				e := b.Put([]byte(ratings[i].ID), j)
				if e != nil {
					log.WithField("error", e).Error("ratings:update:save")
				}
			}
		}

		return nil
	})
}

//ListRatings returns all ratings
func (s *Store) ListRatings() []Rating {
	var (
		ratings []Rating
		err     error
	)

	s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketRatings)
		b.ForEach(func(k, v []byte) error {
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
