// generated by gocipe 8711bda2250e6c55fba92576b5a3047c42d034e6742f4bff5ae1f9e546907463; DO NOT EDIT

package rating

import (
	"database/sql"
	"errors"
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

// Get returns a single Rating from database by primary key
func Get(id string) (*Rating, error) {
	var entity = New()

	if err := crudPreGet(id); err != nil {
		return nil, fmt.Errorf("error executing crudPreGet() in Get(%x) for entity 'Rating': %s", id, err)
	}

	rows, err := db.Query("SELECT t.ratings_count, t.score, t.reaction_summary, t.updated_at FROM ratings t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(entity.RatingsCount, entity.Score, entity.ReactionSummary, entity.UpdatedAt)
		if err != nil {
			return nil, err
		}

	}

	if err = crudPostGet(entity); err != nil {
		return nil, fmt.Errorf("error executing crudPostGet() in Get(%x) for entity 'Rating': %s", id, err)
	}

	return entity, nil
}

// List returns a slice containing Rating records
func List(filters []models.ListFilter) ([]*Rating, error) {
	var (
		list     []*Rating
		segments []string
		values   []interface{}
		err      error
	)

	query := "SELECT t.ratings_count, t.score, t.reaction_summary, t.updated_at FROM ratings"

	if filters, err = crudPreList(filters); err != nil {
		return nil, fmt.Errorf("error executing crudPreList() in List(filters) for entity 'Rating': %s", err)
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
		err := rows.Scan(entity.RatingsCount, entity.Score, entity.ReactionSummary, entity.UpdatedAt)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)

	}

	if list, err = crudPostList(list); err != nil {
		return nil, fmt.Errorf("error executing crudPostList() in List(filters) for entity 'Rating': %s", err)
	}

	return list, nil
}

// Delete deletes a Rating record from database by id primary key
func Delete(id string, tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("DELETE FROM ratings WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Delete(%x) for entity 'Rating': %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%x) for entity 'Rating': %s", id, err)
	}

	if err := crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("Error executing crudPostDelete() in Delete(%x) for entity 'Rating': %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Delete(%x) for 'Rating': %s", id, err)
		}
	}

	return err
}

// Delete deletes a Rating record from database and sets id to nil
func (entity *Rating) Delete(tx *sql.Tx, autocommit bool) error {
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

	stmt, err := tx.Prepare("DELETE FROM ratings WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Rating.Delete() for ID = %x : %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Rating.Delete() for ID = %x : %s", id, err)
	}

	if err = crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostDelete() in Rating.Delete() for ID = %x : %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Rating.Delete() for ID = %x : %s", id, err)
		}
	}

	return err
}

// Save either inserts or updates a Rating record based on whether or not id is nil
func (entity *Rating) Save(tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return errors.New("primary key cannot be nil")
	}
	return entity.Merge(tx, autocommit)
}

// Insert performs an SQL insert for Rating record and update instance with inserted id.
func (entity *Rating) Insert(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}
	*entity.UpdatedAt = time.Now()

	stmt, err := tx.Prepare("INSERT INTO ratings (id, ratings_count, score, reaction_summary, updated_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	if err := crudPreSave("INSERT", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Rating.Insert(): %s", err)
	}

	_, err = stmt.Exec(*entity.ID, *entity.RatingsCount, *entity.Score, *entity.ReactionSummary, *entity.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Rating: %s", err)
	}

	if err := crudPostSave("INSERT", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Rating.Insert(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Rating.Insert(): %s", err)
		}
	}

	return nil
}

// Update Will execute an SQLUpdate Statement for Rating in the database. Prefer using Save instead of Update directly.
func (entity *Rating) Update(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}
	*entity.UpdatedAt = time.Now()

	stmt, err := tx.Prepare("UPDATE ratings SET ratings_count = $2, score = $3, reaction_summary = $4, updated_at = $5 WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreSave("UPDATE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Rating.Update(): %s", err)
	}

	_, err = stmt.Exec(*entity.RatingsCount, *entity.Score, *entity.ReactionSummary, *entity.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Rating.Update(): %s", err)
	}

	if err := crudPostSave("UPDATE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Rating.Update(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Rating.Update(): %s", err)
		}
	}

	return nil
}

// Merge performs an SQL merge for Rating record.
func (entity *Rating) Merge(tx *sql.Tx, autocommit bool) error {
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

	*entity.UpdatedAt = time.Now()

	stmt, err := tx.Prepare(`INSERT INTO ratings (id, ratings_count, score, reaction_summary, updated_at) VALUES ($1, $2, $3, $4, $5) 
	ON CONFLICT ratings_pkey UPDATE ratings SET ratings_count = $2, score = $3, reaction_summary = $4, updated_at = $5 WHERE id = $1`)
	if err != nil {
		return err
	}

	if err := crudPreSave("MERGE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreSave() in Rating.Merge(): %s", err)
	}

	_, err = stmt.Exec(*entity.RatingsCount, *entity.Score, *entity.ReactionSummary, *entity.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Rating: %s", err)
	}

	if err := crudPostSave("MERGE", entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostSave() in Rating.Merge(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Rating.Merge(): %s", err)
		}
	}

	return nil
}