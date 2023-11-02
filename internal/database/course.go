package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()
	createdCourse, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return Course{}, err
	}

	createdCourse.Exec(id, name, description, categoryID)

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}
