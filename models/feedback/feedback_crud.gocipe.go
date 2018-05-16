// generated by gocipe 89440e4dbce197f5b46d60d778de3e709012a9e0bf60a5510bf683dbf5304465; DO NOT EDIT

package feedback

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mscraftsman/devcon-feedback/models"
)

var db *sql.DB

// Inject allows injection of services into the package
func Inject(database *sql.DB) {
	db = database
}

// Get returns a single Feedback from database by primary key
func Get(id int64) (*Feedback, error) {
	var entity = New()

	rows, err := db.Query("SELECT t.id, t.session_id, t.visitor_id, t.reaction_1, t.reaction_2, t.reaction_3, t.reaction_4, t.created_at FROM feedbacks t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(entity.ID, entity.SessionID, entity.VisitorID, entity.Reaction1, entity.Reaction2, entity.Reaction3, entity.Reaction4, entity.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return entity, nil
}

// List returns a slice containing Feedback records
func List(filters []models.ListFilter) ([]*Feedback, error) {
	var (
		list     []*Feedback
		segments []string
		values   []interface{}
		err      error
	)

	query := "SELECT t.id, t.session_id, t.visitor_id, t.reaction_1, t.reaction_2, t.reaction_3, t.reaction_4, t.created_at FROM feedbacks t"

	for i, filter := range filters {
		segments = append(segments, filter.Field+" "+filter.Operation+" $"+strconv.Itoa(i+1))
		values = append(values, filter.Value)
	}

	if len(segments) != 0 {
		query += " WHERE " + strings.Join(segments, " AND ")
	}

	rows, err := db.Query(query+" ORDER BY id ASC", values...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		entity := New()
		err := rows.Scan(entity.ID, entity.SessionID, entity.VisitorID, entity.Reaction1, entity.Reaction2, entity.Reaction3, entity.Reaction4, entity.CreatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)

	}

	return list, nil
}

// Delete deletes a Feedback record from database by id primary key
func Delete(id int64, tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("DELETE FROM feedbacks WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%x) for entity 'Feedback': %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Delete(%x) for 'Feedback': %s", id, err)
		}
	}

	return err
}

// Delete deletes a Feedback record from database and sets id to nil
func (entity *Feedback) Delete(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	id := *entity.ID

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("DELETE FROM feedbacks WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback.Delete() for ID = %x : %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Feedback.Delete() for ID = %x : %s", id, err)
		}
	}

	return err
}

// Save either inserts or updates a Feedback record based on whether or not id is nil
func (entity *Feedback) Save(tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return entity.Insert(tx, autocommit)
	}
	return entity.Update(tx, autocommit)
}

// Insert performs an SQL insert for Feedback record and update instance with inserted id.
func (entity *Feedback) Insert(tx *sql.Tx, autocommit bool) error {
	var (
		id  int64
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}
	*entity.CreatedAt = time.Now()

	stmt, err := tx.Prepare("INSERT INTO feedbacks (session_id, visitor_id, reaction_1, reaction_2, reaction_3, reaction_4, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id")
	if err != nil {
		return err
	}

	err = stmt.QueryRow(*entity.SessionID, *entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt).Scan(&id)
	if err == nil {
		entity.ID = &id
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback: %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Feedback.Insert(): %s", err)
		}
	}

	return nil
}

// Update Will execute an SQLUpdate Statement for Feedback in the database. Prefer using Save instead of Update directly.
func (entity *Feedback) Update(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("UPDATE feedbacks SET session_id = $1, visitor_id = $2, reaction_1 = $3, reaction_2 = $4, reaction_3 = $5, reaction_4 = $6, created_at = $7 WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(*entity.SessionID, *entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback.Update(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Feedback.Update(): %s", err)
		}
	}

	return nil
}

// Merge performs an SQL merge for Feedback record.
func (entity *Feedback) Merge(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	if entity.ID == nil {
		return entity.Insert(tx, autocommit)
	}

	*entity.CreatedAt = time.Now()

	stmt, err := tx.Prepare(`INSERT INTO feedbacks (id, session_id, visitor_id, reaction_1, reaction_2, reaction_3, reaction_4, created_at) VALUES ($1, $1, $2, $3, $4, $5, $6, $7) 
	ON CONFLICT (id) DO UPDATE SET session_id = $1, visitor_id = $2, reaction_1 = $3, reaction_2 = $4, reaction_3 = $5, reaction_4 = $6, created_at = $7`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(*entity.ID, *entity.SessionID, *entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback: %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Feedback.Merge(): %s", err)
		}
	}

	return nil
}
