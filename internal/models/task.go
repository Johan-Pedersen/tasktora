package models

import (
	"database/sql"
	"time"
)

type Task struct {
	Id       int
	Title    string
	Note     string
	Created  time.Time
	ParentId sql.NullInt64
	Level    int
}

type TaskModel struct {
	DB *sql.DB
}

type ITaskModel interface {
	Get(id int) (*Task, error)
	GetAll() ([]*Task, error)
}

func (tm TaskModel) Get(id int) (*Task, error) {
	stmt := `SELECT *
  FROM tasks
  WHERE id = ?
  `

	row := tm.DB.QueryRow(stmt, id)

	task := &Task{}

	err := row.Scan(&task.Id, &task.Title, &task.Note, &task.Created, &task.ParentId, &task.Level)
	if err != nil {
		return nil, err
	}
	// If everything went OK then return the Snippet object.
	return task, nil
}

// func (tm *TaskModel) GetChildren()(*[]Task, error){
//
//	stmt := `SELECT id, title, note, created, parent_id FROM tasks
//
// WHERE id = ?
//
//	  `
//
//		row := tm.DB.Query(stmt, id)
//
//		task := &Task{}
//
//		err := row.Scan(&task.Id, &task.Title, &task.Note, &task.Created, &task.ParentId)
//		if err != nil {
//			// If the query returns no rows, then row.Scan() will return a
//			// sql.ErrNoRows error. We use the errors.Is() function check for that
//			// error specifically, and return our own ErrNoRecord error
//			// instead (we'll create this in a moment).
//			if errors.Is(err, sql.ErrNoRows) {
//				return nil, ErrNoRecord
//			} else {
//				return nil, err
//			}
//		}
//		// If everything went OK then return the Snippet object.
//		return task, nil
//	}

func (tm TaskModel) GetAll() ([]*Task, error) {
	stmt := `SELECT *
  FROM tasks 
  `

	rows, err := tm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	tasks := []*Task{}

	defer rows.Close()

	for rows.Next() {
		task := &Task{}

		err = rows.Scan(&task.Id, &task.Title, &task.Note, &task.Created, &task.ParentId, &task.Level)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippet object.
	return tasks, nil
}
