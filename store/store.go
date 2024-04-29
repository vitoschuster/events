package store

import (
	"database/sql"

	"github.com/vitoschuster/events/types"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	CreateEvent(event *types.Event) (*types.Event, error)
	GetEvents() ([]types.Event, error)
	DeleteEvent(id string) error
	FindEventsByLocationNameOrDesc(search string) ([]types.Event, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) DeleteEvent(id string) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id = ?", id)
	return err
}

func (s *Storage) CreateEvent(c *types.Event) (*types.Event, error) {
	row, err := s.db.Exec("INSERT INTO events (location, event_time, name, description, price, imageURL) VALUES (?, ?, ?, ?, ?)", c.Location, c.EventTime, c.Name, c.Description, c.Price, c.ImageURL)
	if err != nil {
		return nil, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	c.ID = int(id)

	return c, nil
}

func (s *Storage) GetEvents() ([]types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []types.Event
	for rows.Next() {
		event, err := scanEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (s *Storage) FindEventsByLocationNameOrDesc(search string) ([]types.Event, error) {
	rows, err := s.db.Query("SELECT * FROM events WHERE location LIKE ? OR name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []types.Event
	for rows.Next() {
		event, err := scanEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func scanEvent(row *sql.Rows) (types.Event, error) {
	var event types.Event
	err := row.Scan(&event.ID, &event.Location, &event.EventTime, &event.Name, &event.Description, &event.Price, &event.ImageURL)
	return event, err
}
