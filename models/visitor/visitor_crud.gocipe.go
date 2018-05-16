// generated by gocipe 89440e4dbce197f5b46d60d778de3e709012a9e0bf60a5510bf683dbf5304465; DO NOT EDIT

package visitor

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mscraftsman/devcon-feedback/models"
)

var db *sql.DB

// Inject allows injection of services into the package
func Inject(database *sql.DB) {
	db = database
}

// Get returns a single Visitor from database by primary key
func Get(id string) (*Visitor, error) {
	var entity = New()

	rows, err := db.Query("SELECT t.id, t.name, t.photo_link, t.status FROM visitors t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(entity.ID, entity.Name, entity.PhotoLink, entity.Status)
		if err != nil {
			return nil, err
		}

	}

	return entity, nil
}

// List returns a slice containing Visitor records
func List(filters []models.ListFilter) ([]*Visitor, error) {
	var (
		list     []*Visitor
		segments []string
		values   []interface{}
		err      error
	)

	query := "SELECT t.id, t.name, t.photo_link, t.status FROM visitors t"

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
		err := rows.Scan(entity.ID, entity.Name, entity.PhotoLink, entity.Status)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)

	}

	return list, nil
}

// Delete deletes a Visitor record from database by id primary key
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

	stmt, err := tx.Prepare("DELETE FROM visitors WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%x) for entity 'Visitor': %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Delete(%x) for 'Visitor': %s", id, err)
		}
	}

	return err
}

// Delete deletes a Visitor record from database and sets id to nil
func (entity *Visitor) Delete(tx *sql.Tx, autocommit bool) error {
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

	stmt, err := tx.Prepare("DELETE FROM visitors WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor.Delete() for ID = %x : %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Delete() for ID = %x : %s", id, err)
		}
	}

	return err
}

// Save either inserts or updates a Visitor record based on whether or not id is nil
func (entity *Visitor) Save(tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return errors.New("primary key cannot be nil")
	}
	return entity.Merge(tx, autocommit)
}

// Insert performs an SQL insert for Visitor record and update instance with inserted id.
func (entity *Visitor) Insert(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("INSERT INTO visitors (id, name, photo_link, status) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(*entity.ID, *entity.Name, *entity.PhotoLink, *entity.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor: %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Insert(): %s", err)
		}
	}

	return nil
}

// Update Will execute an SQLUpdate Statement for Visitor in the database. Prefer using Save instead of Update directly.
func (entity *Visitor) Update(tx *sql.Tx, autocommit bool) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}

	stmt, err := tx.Prepare("UPDATE visitors SET name = $2, photo_link = $3, status = $4 WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(*entity.Name, *entity.PhotoLink, *entity.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor.Update(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Update(): %s", err)
		}
	}

	return nil
}

// Merge performs an SQL merge for Visitor record.
func (entity *Visitor) Merge(tx *sql.Tx, autocommit bool) error {
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

	stmt, err := tx.Prepare(`INSERT INTO visitors (id, name, photo_link, status) VALUES ($1, $2, $3, $4) 
	ON CONFLICT (id) DO UPDATE SET name = $2, photo_link = $3, status = $4`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(*entity.ID, *entity.Name, *entity.PhotoLink, *entity.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor: %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Merge(): %s", err)
		}
	}

	return nil
}
