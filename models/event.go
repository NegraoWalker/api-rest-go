package models

import (
	"api-rest/db"
	"time"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
    INSERT INTO events(name, description, location, dateTime, user_id)
    VALUES ($1, $2, $3, $4, $5) RETURNING id
    `
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"

	row := db.DB.QueryRow(query, id)

	var event Event
	error := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if error != nil {
		return nil, error
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (event Event) Update() error {
	query := `
    UPDATE events
    SET name = $1, description = $2, location = $3, dateTime = $4
    WHERE id = $5
    `
	stmt, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer stmt.Close()

	_, error = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return error
}
