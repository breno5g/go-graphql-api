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

func (c *Course) List() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return []Course{}, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course

		rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses WHERE category_id = ?", categoryID)
	if err != nil {
		return []Course{}, err
	}

	var courses []Course
	for rows.Next() {
		var course Course

		rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)

		courses = append(courses, course)
	}

	return courses, nil
}
