package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

var events = []Event{}

func (e Event) Save(db *sql.DB) error {
	// Saving to Database
	fmt.Println(e.UserID)
	query := "INSERT INTO events (name, description, location, dateTime, user_id) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	return err
}

func GetAllEvents(db *sql.DB) ([]Event, error) {
	query := `SELECT * from events`
	fmt.Println("DB Details", db)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		// Gets the data and maps to the field. Order should be same as was a create table is made
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64, db *sql.DB) (*Event, error) {
	query := "SELECT * FROM EVENTS WHERE id = $1"
	row := db.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(event Event, db *sql.DB) error {
	query := `UPDATE EVENTS
	SET 
	name = $1, description = $2, location = $3, dateTime = $4
	WHERE ID = $5
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return err
}

func DeleteEvent(eventId int64, db *sql.DB) error {
	query := `DELETE FROM EVENTS WHERE id=$1`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(eventId)
	if err != nil {
		return err
	}
	return err
}
