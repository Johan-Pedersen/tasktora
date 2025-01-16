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

func (tm *TaskModel) GetTaskAndSubTasks(id int) ([]*Task, error) {
	stmt := `WITH RECURSIVE cte AS (
    (SELECT t.id, t.title, t.note, t.created, t.parent_id, 1 as level
    FROM tasks t
    WHERE t.id = ?)
  union all
    (SELECT this.id, this.title, this.note, this.created, this.parent_id, this.level
    FROM cte prior
    INNER JOIN tasks this ON 
      this.parent_id = prior.id)
)

SELECT e.id, e.title, e.note, e.created, e.parent_id, e.level
FROM cte e
  `
	rows, err := tm.DB.Query(stmt, id)
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
