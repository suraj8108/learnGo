package models

import "database/sql"

type RegisterEvent struct {
	EventId int64 `binding:"required"`
	UserId  int64 `binding:"required"`
}

func (r RegisterEvent) RegisterUserForEvent(db *sql.DB) error {
	query := `
		INSERT INTO registration (event_id, user_id ) 
		VALUES ($1, $2)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(r.EventId, r.UserId)
	if err != nil {
		return err
	}
	return err

}

func (r RegisterEvent) CancelRegistration(db *sql.DB) error {
	query := `
		DELETE FROM Registration where event_id = $1 and user_id = $2 
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(r.EventId, r.UserId)
	if err != nil {
		return err
	}
	return err

}

func (r RegisterEvent) GetRegistrationOfUserForEvent(db *sql.DB) (*RegisterEvent, error) {
	query := `
		SELECT event_id, user_id from registration where event_id = $1 and user_id = $2 
	`

	row := db.QueryRow(query, r.EventId, r.UserId)
	var dbDetails RegisterEvent
	err := row.Scan(&dbDetails.EventId, &dbDetails.UserId)
	if err != nil {
		return nil, err
	}

	return &dbDetails, nil

}
