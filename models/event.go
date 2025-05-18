package models

import (
	"fmt"
	"time"

	"github.com/RevanthGovindan/event-booking/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
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
	fmt.Println("Inserted ID:", id)
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
