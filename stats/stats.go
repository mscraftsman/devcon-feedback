package stats

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/util"
)

var db *sql.DB

type sessionSummary struct {
	SessionID string `json:"session_id"`
	Title     string `json:"title"`
	Speakers  string `json:"speakers"`
	Reaction1 int    `json:"reaction_1"`
	Reaction2 int    `json:"reaction_2"`
}

type voterSummary struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Votes int    `json:"votes"`
}

// Inject dependencies
func Inject(d *sql.DB) {
	db = d
}

// Summary is an endpoint for stats summary
func Summary(w http.ResponseWriter, r *http.Request) {
	var (
		sessions []sessionSummary
		voters   []voterSummary
		wg       sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		sessions = topSessions()
		wg.Done()
	}()

	go func() {
		voters = topVoters()
		wg.Done()
	}()

	wg.Wait()

	util.JSONOutputResponse(w, struct {
		Sessions []sessionSummary `json:"sessions"`
		Voters   []voterSummary   `json:"voters"`
	}{sessions, voters})
}

func topSessions() []sessionSummary {
	var sessions []sessionSummary
	rows, err := db.Query("select session_id, sum(cast(reaction_1 as int)) r1, sum(cast(reaction_2 as int)) r2 from feedbacks group by session_id order by r1 desc, r2 desc")
	if err != nil {
		log.Println(err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var s sessionSummary
		err := rows.Scan(&s.SessionID, &s.Reaction1, &s.Reaction2)
		if err != nil {
			return nil
		}
		sess := sessionize.Get(s.SessionID)

		s.Title = sess.Title
		s.Speakers = strings.Join(sess.Speakers, ", ")

		sessions = append(sessions, s)
	}

	return sessions
}

func topVoters() []voterSummary {
	var voters []voterSummary
	rows, err := db.Query(`
select v.id, v.name, v.photo_link, f.t from (
	select count(visitor_id) t, visitor_id from feedbacks group by visitor_id order by t desc
) f inner join visitors v on ( f.visitor_id = v.id) order by f.t desc
`)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var v voterSummary
		err := rows.Scan(&v.ID, &v.Name, &v.Photo, &v.Votes)
		if err != nil {
			return nil
		}

		voters = append(voters, v)
	}

	return voters
}
