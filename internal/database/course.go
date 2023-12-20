package database

// import (
// 	"database/sql"
// 	"github.com/google/uuid"
// )

// type Course struct {
// 	db          *sql.DB
// 	ID          string
// 	Name        string
// 	Description string
// 	Category    *Category
// }

// func NewCourse(db *sql.DB) *Course {
// 	return &Course{
// 		db: db,
// 	}
// }

// func (course *Course) Create(name string, description string, category Category) (Course, error) {
// 	id := uuid.New().String()

// 	_, err := course.db.Exec(
// 		"INSERT INTO COURSES (id, name, description, cours) VALUES	($1, $2, $3, $4)",
// 		id, name, description,
// 	)

// 	if err != nil {
// 		return Course{}, err
// 	}
// 	return Course{ID: id, Name: name, Description: description, Category: &category}, nil
// }
