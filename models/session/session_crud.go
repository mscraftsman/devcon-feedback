package session

import "database/sql"

var db *sql.DB

// Inject allows injection of services into the package
func Inject(database *sql.DB) {
	db = database
}

// Get returns a single Session from database by primary key
func Get(id int64) (*Session, error) {
	var entity = New()
	
    if err := crudPreGet(id); err != nil {
		return nil, fmt.Errorf("error executing crudPreGet() in Get(%d) for entity 'Session': %s", id, err)
	}
    
	rows, err := db.Query("SELECT t.id, t.title, t.description, t.level, t.language, t.format, t.room, t.speakers, t.ratings_count, t.score, t.reaction_summary, t.start_at, t.end_at, t.status FROM  t WHERE id = $1 ORDER BY t.id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		

		err := rows.Scan(entity.Id, entity.Title, entity.Description, entity.Level, entity.Language, entity.Format, entity.Room, entity.Speakers, entity.RatingsCount, entity.Score, entity.ReactionSummary, entity.StartAt, entity.EndAt, entity.Status)
		if err != nil {
			return nil, err
		} 
		
		
	}
	
	if err = crudPostGet(entity); err != nil {
		return nil, fmt.Errorf("error executing crudPostGet() in Get(%d) for entity 'Session': %s", id, err)
	}
	

	return entity, nil
}

// List returns a slice containing Session records
func List(filters []models.ListFilter) ([]*Session, error) {
	var (
		list     []*Session
		segments []string
		values   []interface{}
		err      error
	)

	query := "SELECT t.id, t.title, t.description, t.level, t.language, t.format, t.room, t.speakers, t.ratings_count, t.score, t.reaction_summary, t.start_at, t.end_at, t.status FROM "
	
    if filters, err = crudPreList(filters); err != nil {
		return nil, fmt.Errorf("error executing crudPreList() in List(filters) for entity 'Session': %s", err)
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
		err := rows.Scan(entity.Id, entity.Title, entity.Description, entity.Level, entity.Language, entity.Format, entity.Room, entity.Speakers, entity.RatingsCount, entity.Score, entity.ReactionSummary, entity.StartAt, entity.EndAt, entity.Status)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
		
	}
	

	
	if list, err = crudPostList(list); err != nil {
		return nil, fmt.Errorf("error executing crudPostList() in List(filters) for entity 'Session': %s", err)
	}
	
	return list, nil
}



// Delete deletes a Session record from database by id primary key
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

	stmt, err := tx.Prepare("DELETE FROM  WHERE id = $1")
	if err != nil {
		return err
	}
	
	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Delete(%d) for entity 'Session': %s", id, err)
	}
	
	
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Delete(%d) for entity 'Session': %s", id, err)
	}
	
	if err := crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("Error executing crudPostDelete() in Delete(%d) for entity 'Session': %s", id, err)
	}
	
	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Delete(%d) for 'Session': %s", id, err)
		}
	}

	return err
}

// Delete deletes a Session record from database and sets id to nil
func (entity *Session) Delete(tx *sql.Tx, autocommit bool) error {
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

	stmt, err := tx.Prepare("DELETE FROM  WHERE id = $1")
	if err != nil {
		return err
	}
	
	if err := crudPreDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreDelete() in Session.Delete() for ID = %d : %s", id, err)
	}
	
	
	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = nil
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Session.Delete() for ID = %d : %s", id, err)
	}
	
	if err = crudPostDelete(id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostDelete() in Session.Delete() for ID = %d : %s", id, err)
	}
	
	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Session.Delete() for ID = %d : %s", id, err)
		}
	}

	return err
}

// Save either inserts or updates a Session record based on whether or not id is nil
func (entity *Session) Save(tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return entity.Insert(tx, autocommit)
	}
	return entity.Update(tx, autocommit)
}

// Insert performs an SQL insert for Session record and update instance with inserted id.
// Prefer using Save rather than Insert directly.
func (entity *Session) Insert(tx *sql.Tx, autocommit bool) error {
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
	
	stmt, err := tx.Prepare("INSERT INTO  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id")
	if err != nil {
		return err
	}
	
    if err := crudPreCreate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPreCreate() in Session.Insert(): %s", err)
	}
    
	err = stmt.QueryRow(*entity.Id, *entity.Title, *entity.Description, *entity.Level, *entity.Language, *entity.Format, *entity.Room, *entity.Speakers, *entity.RatingsCount, *entity.Score, *entity.ReactionSummary, *entity.StartAt, *entity.EndAt, *entity.Status).Scan(&id)
	if err == nil {
		entity.ID = &id
	} else {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Session: %s", err)
	}
	
	
	if err := crudPostCreate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostCreate() in Session.Insert(): %s", err)
	}
	
	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Session.Insert(): %s", err)
		}
	}

	return nil
}

// Update Will execute an SQLUpdate Statement for Session in the database. Prefer using Save instead of Update directly.
func (entity *Session) Update(tx *sql.Tx, autocommit bool) error {
	var (
		err error
		
	)

	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return err
		}
	}
	
	stmt, err := tx.Prepare("UPDATE  SET id = $1, title = $2, description = $3, level = $4, language = $5, format = $6, room = $7, speakers = $8, ratings_count = $9, score = $10, reaction_summary = $11, start_at = $12, end_at = $13, status = $14 WHERE id = $1")
	if err != nil {
		return err
	}

	
    if err := crudPreUpdate(entity, tx); err != nil {
		tx.Rollback()
        return fmt.Errorf("error executing crudPreUpdate() in Session.Update(): %s", err)
	}
    
	_, err = stmt.Exec(*entity.Id, *entity.Title, *entity.Description, *entity.Level, *entity.Language, *entity.Format, *entity.Room, *entity.Speakers, *entity.RatingsCount, *entity.Score, *entity.ReactionSummary, *entity.StartAt, *entity.EndAt, *entity.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing transaction statement in Session.Update(): %s", err)
	}
	
	
	if err := crudPostUpdate(entity, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing crudPostUpdate() in Session.Update(): %s", err)
	}
	
	if autocommit {
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("error committing transaction in Session.Update(): %s", err)
		}
	}

	return nil
}