package models

import (
	"database/sql"
	"log"
	"tasktora/internal/assert"
	"testing"
	"time"
)

func TestGetTask(t *testing.T) {
	created, err := time.Parse("2006-01-02 15:04:05", "2025-01-01 10:00:00")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name    string
		idInput int
		want    *Task
	}{
		{
			name:    "Valid ID, parent",
			idInput: 1,
			want: &Task{
				Title:    "fugl 1",
				Note:     "Synger",
				Id:       1,
				Created:  created,
				ParentId: sql.NullInt64{Valid: false},
				Level:    1,
			},
		},
		{
			name:    "Valid ID, child",
			idInput: 3,
			want: &Task{
				Title:    "fugl 3",
				Note:     "Gaar",
				Id:       3,
				ParentId: sql.NullInt64{Int64: 1, Valid: true},
				Created:  created,
				Level:    2,
			},
		},
		{
			name:    "Zero ID",
			idInput: 0,
			want:    nil,
		},
		{
			name:    "Non-existent ID",
			idInput: 9,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := TaskModel{db}

			task, err := m.GetTask(tt.idInput)
			assert.Equal(t, task, tt.want)
			if task == nil {
				assert.ErrNoRows(t, err)
			}
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	created, err := time.Parse("2006-01-02 15:04:05", "2025-01-01 10:00:00")
	if err != nil {
		log.Fatal(err)
	}
	t.Run("GetAllTasks())", func(t *testing.T) {
		db := newTestDB(t)

		m := TaskModel{db}

		tasks, err := m.GetAllTasks()

		expectedTasks := []*Task{
			{
				Title:   "fugl 1",
				Note:    "Synger",
				Created: created,
				Id:      1,
				Level:   1,
			},
			{
				Title:   "fugl 2",
				Note:    "flyver",
				Created: created,
				Id:      2,
				Level:   1,
			},
			{
				Title:    "fugl 3",
				Note:     "Gaar",
				Created:  created,
				Id:       3,
				ParentId: sql.NullInt64{Int64: 1, Valid: true},
				Level:    2,
			},
			{
				Title:    "fugl 4",
				Note:     "danser",
				Created:  created,
				Id:       4,
				ParentId: sql.NullInt64{Int64: 3, Valid: true},
				Level:    3,
			},
			{
				Title:    "fugl 5",
				Note:     "sover",
				Created:  created,
				Id:       5,
				ParentId: sql.NullInt64{Int64: 1, Valid: true},
				Level:    2,
			},
		}

		assert.Equal(t, tasks, expectedTasks)
		if tasks == nil {
			assert.ErrNoRows(t, err)
		}
	})
}

func TestGetTaskAndSubTasks(t *testing.T) {
	created, err := time.Parse("2006-01-02 15:04:05", "2025-01-01 10:00:00")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name    string
		idInput int
		want    []*Task
	}{
		{
			name:    "Parent children",
			idInput: 1,
			want: []*Task{
				{
					Title:    "fugl 1",
					Note:     "Synger",
					Id:       1,
					Created:  created,
					ParentId: sql.NullInt64{Int64: 0, Valid: false},
					Level:    1,
				},
				{
					Title:    "fugl 3",
					Note:     "Gaar",
					Id:       3,
					ParentId: sql.NullInt64{Int64: 1, Valid: true},
					Created:  created,
					Level:    2,
				},
				{
					Title:    "fugl 5",
					Note:     "sover",
					Id:       5,
					ParentId: sql.NullInt64{Int64: 1, Valid: true},
					Created:  created,
					Level:    2,
				},
				{
					Title:    "fugl 4",
					Note:     "danser",
					Id:       4,
					ParentId: sql.NullInt64{Int64: 3, Valid: true},
					Created:  created,
					Level:    3,
				},
			},
		},
		{
			name:    "No children",
			idInput: 2,
			want: []*Task{
				{
					Title:    "fugl 2",
					Note:     "flyver",
					Id:       2,
					ParentId: sql.NullInt64{Valid: false},
					Created:  created,
					Level:    1,
				},
			},
		},
		{
			name:    "Non-existent ID",
			idInput: 9,
			want:    []*Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := TaskModel{db}

			tasks, _ := m.GetTaskAndSubTasks(tt.idInput)

			assert.Equal(t, tasks, tt.want)
		})
	}
}

func TestInsertTask(t *testing.T) {
	tests := []struct {
		name      string
		taskInput Task
		want      *Task
	}{
		{
			name: "insert parent",
			taskInput: Task{
				Title:    "fugl 9",
				Note:     "lober",
				ParentId: sql.NullInt64{Valid: false},
				Level:    1,
			},
			want: &Task{
				Title:    "fugl 9",
				Note:     "lober",
				Id:       6,
				ParentId: sql.NullInt64{Valid: false},
				Level:    1,
			},
		},
		{
			name: "insert child",
			taskInput: Task{
				Title:    "fugl 10",
				Note:     "Taenker",
				ParentId: sql.NullInt64{Int64: 6, Valid: true},
				Level:    2,
			},
			want: &Task{
				Title:    "fugl 10",
				Note:     "Taenker",
				Id:       6,
				ParentId: sql.NullInt64{Int64: 6, Valid: true},
				Level:    2,
			},
		},
		{
			name: "insert child, but parent doesnt exist",
			taskInput: Task{
				Title:    "fugl 10",
				Note:     "Taenker",
				ParentId: sql.NullInt64{Int64: 7, Valid: true},
				Level:    2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := TaskModel{db}

			ti := tt.taskInput
			id, err := m.InsertTask(ti.Title, ti.Note, ti.ParentId, ti.Level)
			if id == 0 {
				assert.ErrNoRows(t, err)
			} else {

				task, _ := m.GetTask(id)
				want := tt.want
				assert.Equal(t, task.Title, want.Title)
				assert.Equal(t, task.Note, want.Note)
				assert.Equal(t, task.Level, want.Level)
				assert.Equal(t, task.ParentId, want.ParentId)
				assert.Equal(t, task.Id, want.Id)
			}
		})
	}
}
