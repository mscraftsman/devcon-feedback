package visitor

import (
	"database/sql"
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
func Get(id int64) (*Visitor, error) {
	var entity = New()

	if err := crudPreGet(id); err != nil {
		return nil, fmt.Errorf("error executing crudPreGet() in Get(%d) for entity 'Visitor': %s", id, err)
	}

	rows, err := db.Query("SELECT t.name, t.meetup_id, t.photo_link FROM visitors t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(entity.Name, entity.MeetupID, entity.PhotoLink)
		if err != nil {
			return nil, err
		}

	}

	if err = crudPostGet(entity); err != nil {
		return nil, fmt.Errorf("error executing crudPostGet() in Get(%d) for entity 'Visitor': %s", id, err)
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

	query := "SELECT t.name, t.meetup_id, t.photo_link FROM visitors"

	if filters, err = crudPreList(filters); err != nil {
		return nil, fmt.Errorf("error executing crudPreList() in List(filters) for entity 'Visitor': %s", err)
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
		err := rows.Scan(entity.Name, entity.MeetupID, entity.PhotoLink)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)

	}

	if list, err = crudPostList(list); err != nil {
		return nil, fmt.Errorf("error executing crudPostList() in List(filters) for entity 'Visitor': %s", err)
	}

	return list, nil
}

// Delete deletes a Visitor record from database by id primary key
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

	stmt, err := tx.Prepare("DELETE FROM visitors WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Delete(%d) for entity 'Visitor': %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%d) for entity 'Visitor': %s", id, err)
	}

	if err := crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("Error executing crudPostDelete() in Delete(%d) for entity 'Visitor': %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Delete(%d) for 'Visitor': %s", id, err)
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

	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Visitor.Delete() for ID = %d : %s", id, err)
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor.Delete() for ID = %d : %s", id, err)
	}

	if err = crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostDelete() in Visitor.Delete() for ID = %d : %s", id, err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Delete() for ID = %d : %s", id, err)
		}
	}

	return err
}

// Save either inserts or updates a Visitor record based on whether or not id is nil
func (entity *Visitor) Save(tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return entity.Insert(tx, autocommit)
	}
	return entity.Update(tx, autocommit)
}

// Insert performs an SQL insert for Visitor record and update instance with inserted id.
// Prefer using Save rather than Insert directly.
func (entity *Visitor) Insert(tx *sql.Tx, autocommit bool) error {
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

	stmt, err := tx.Prepare("INSERT INTO visitors ($1, $2, $3) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}

	if err := crudPreCreate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreCreate() in Visitor.Insert(): %s", err)
	}

	err = stmt.QueryRow(*entity.Name, *entity.MeetupID, *entity.PhotoLink).Scan(&id)
	if err == nil {
		entity.ID = &id
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor: %s", err)
	}

	if err := crudPostCreate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostCreate() in Visitor.Insert(): %s", err)
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

	stmt, err := tx.Prepare("UPDATE visitors SET name = $1, meetup_id = $2, photo_link = $3 WHERE id = $1")
	if err != nil {
		return err
	}

	if err := crudPreUpdate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreUpdate() in Visitor.Update(): %s", err)
	}

	_, err = stmt.Exec(*entity.Name, *entity.MeetupID, *entity.PhotoLink)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Visitor.Update(): %s", err)
	}

	if err := crudPostUpdate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostUpdate() in Visitor.Update(): %s", err)
	}

	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Visitor.Update(): %s", err)
		}
	}

	return nil
}
