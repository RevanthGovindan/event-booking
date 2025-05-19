package models

import (
	"time"

	"github.com/RevanthGovindan/event-booking/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name,description,location,dateTime,user_id) VALUES(?,?,?,?,?)`
	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	query := `update events SET name =?,description=?,location=?,dateTime=? where id = ?`
	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if count > 0 {
		return nil
	}
	return err
}

func (e Event) Delete() error {
	query := `delete from events where id = ?`
	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.ID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if count > 0 {
		return nil
	}
	return err
}

func GetAllEvents() ([]Event, error) {
	var result []Event = make([]Event, 0)
	query := `select * from events`
	rows, err := db.Db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return result, err
		}
		result = append(result, event)
	}
	return result, nil
}

func GetEventById(id int64) (Event, error) {
	query := `select * from events where id = ?`
	rows := db.Db.QueryRow(query, id)

	var event Event
	err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return event, err
	}
	return event, nil
}
