// generated by gocipe 2e5ed5d4d943c424d9e2d746f06bc60d148cc25203f3693488ed10ba23c0d12a; DO NOT EDIT

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

	if err := crudPreGet(id); err != nil {
		return nil, fmt.Errorf("error executing crudPreGet() in Get(%x) for entity 'Feedback': %s", id, err)
	}

	rows, err := db.Query("SELECT t.visitor_id, t.reaction_1, t.reaction_2, t.reaction_3, t.reaction_4, t.created_at FROM feedbacks t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(entity.VisitorID, entity.Reaction1, entity.Reaction2, entity.Reaction3, entity.Reaction4, entity.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	if err = crudPostGet(entity); err != nil {
		return nil, fmt.Errorf("error executing crudPostGet() in Get(%x) for entity 'Feedback': %s", id, err)
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

	query := "SELECT t.visitor_id, t.reaction_1, t.reaction_2, t.reaction_3, t.reaction_4, t.created_at FROM feedbacks"

	if filters, err = crudPreList(filters); err != nil {
		return nil, fmt.Errorf("error executing crudPreList() in List(filters) for entity 'Feedback': %s", err)
	}

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
		err := rows.Scan(entity.VisitorID, entity.Reaction1, entity.Reaction2, entity.Reaction3, entity.Reaction4, entity.CreatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)

	}

	if list, err = crudPostList(list); err != nil {
		return nil, fmt.Errorf("error executing crudPostList() in List(filters) for entity 'Feedback': %s", err)
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

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Delete(%x) for entity 'Feedback': %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%x) for entity 'Feedback': %s", id, err)
	}

	if err := crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("Error executing crudPostDelete() in Delete(%x) for entity 'Feedback': %s", id, err)
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

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Feedback.Delete() for ID = %x : %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback.Delete() for ID = %x : %s", id, err)
	}

	if err = crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostDelete() in Feedback.Delete() for ID = %x : %s", id, err)
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

	stmt, err := tx.Prepare("INSERT INTO feedbacks (visitor_id, reaction_1, reaction_2, reaction_3, reaction_4, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return err
	}

	if err := crudPreSave("INSERT", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Feedback.Insert(): %s", err)
	}

	err = stmt.QueryRow(*entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt).Scan(&id)
	if err == nil {
		entity.ID = &id
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback: %s", err)
	}

	if err := crudPostSave("INSERT", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Feedback.Insert(): %s", err)
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

	stmt, err := tx.Prepare("UPDATE feedbacks SET visitor_id = $1, reaction_1 = $2, reaction_2 = $3, reaction_3 = $4, reaction_4 = $5, created_at = $6 WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreSave("UPDATE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Feedback.Update(): %s", err)
	}

	_, err = stmt.Exec(*entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback.Update(): %s", err)
	}

	if err := crudPostSave("UPDATE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Feedback.Update(): %s", err)
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

	stmt, err := tx.Prepare(`INSERT INTO feedbacks (id, visitor_id, reaction_1, reaction_2, reaction_3, reaction_4, created_at) VALUES ($1, $1, $2, $3, $4, $5, $6) 
	ON CONFLICT (id) DO UPDATE SET visitor_id = $1, reaction_1 = $2, reaction_2 = $3, reaction_3 = $4, reaction_4 = $5, created_at = $6`)
	if err != nil {
		return err
	}

	if err := crudPreSave("MERGE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Feedback.Merge(): %s", err)
	}

	_, err = stmt.Exec(*entity.ID, *entity.VisitorID, *entity.Reaction1, *entity.Reaction2, *entity.Reaction3, *entity.Reaction4, *entity.CreatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Feedback: %s", err)
	}

	if err := crudPostSave("MERGE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Feedback.Merge(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Feedback.Merge(): %s", err)
		}
	}

	return nil
}
