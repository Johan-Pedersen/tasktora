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
	GetTask(id int) (*Task, error)
	GetAllTasks() ([]*Task, error)
	GetTaskAndSubTasks(id int) ([]*Task, error)
	InsertTask(title, note string, parentId sql.NullInt64, level int) (int, error)
	UpdateTask(id int, title, note string, parentId sql.NullInt64, level int) (int, error)
	UpdateTaskTitle(id int, title string) (int, error)
	UpdateTaskNote(id int, note string) (int, error)
	UpdateTaskParentId(id int, parentId sql.NullInt64) (int, error)
	UpdateTaskLvl(id int, level int) (int, error)
}

func (tm TaskModel) GetTask(id int) (*Task, error) {
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

func (tm TaskModel) GetTaskAndSubTasks(id int) ([]*Task, error) {
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

func (tm TaskModel) GetAllTasks() ([]*Task, error) {
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

func (tm TaskModel) InsertTask(title, note string, parentId sql.NullInt64, level int) (int, error) {
	// Verify parent exists
	_, err := tm.GetTask(int(parentId.Int64))
	if err != nil {
		return 0, err
	}

	insStmt := `INSERT INTO tasks (title, note, created, parent_id, level)
VALUES(?, ?, UTC_TIMESTAMP(),? , ?)`

	result, err := tm.DB.Exec(insStmt, title, note, parentId, level)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm TaskModel) UpdateTask(id int, title, note string, parentId sql.NullInt64, level int) (int, error) {
	// Verify parent exists

	stmt := `UPDATE tasks 
  SET title = ?, note = ?, parent_id = ?, level = ?
  WHERE id = ?`

	_, err := tm.DB.Exec(stmt, title, note, parentId, level, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm TaskModel) UpdateTaskTitle(id int, title string) (int, error) {
	// Verify parent exists

	stmt := `UPDATE tasks 
  SET title = ?
  WHERE id = ?`

	_, err := tm.DB.Exec(stmt, title, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm TaskModel) UpdateTaskNote(id int, note string) (int, error) {
	// Verify parent exists

	stmt := `UPDATE tasks 
  SET note = ?
  WHERE id = ?`

	_, err := tm.DB.Exec(stmt, note)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm TaskModel) UpdateTaskParent(id int, parentId sql.NullInt64) (int, error) {
	// Verify parent exists

	stmt := `UPDATE tasks 
  SET parent_id = ?
  WHERE id = ?`

	_, err := tm.DB.Exec(stmt, parentId, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm TaskModel) UpdateTaskLvl(id int, level int) (int, error) {
	// Verify parent exists

	stmt := `UPDATE tasks 
  SET level = ?
  WHERE id = ?`

	_, err := tm.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
