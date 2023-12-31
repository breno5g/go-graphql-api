package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	createdCategory, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)")
	if err != nil {
		return Category{}, err
	}

	createdCategory.Exec(id, name, description)

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) List() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM categories")
	if err != nil {
		return []Category{}, err
	}

	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category

		rows.Scan(&category.ID, &category.Name, &category.Description)

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var category Category

	row := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID)

	err := row.Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return Category{}, err
	}

	return category, nil
}
