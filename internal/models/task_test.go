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
			// Call the newTestDB() helper function to get a connection pool to
			// our test database. Calling this here -- inside t.Run() -- means
			// that fresh database tables and data will be set up and torn down
			// for each sub-test.
			db := newTestDB(t)
			// Create a new instance of the UserModel.
			m := TaskModel{db}
			// Call the UserModel.Exists() method and check that the return
			// value and error match the expected values for the sub-test.
			task, err := m.Get(tt.idInput)
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
		// Call the newTestDB() helper function to get a connection pool to
		// our test database. Calling this here -- inside t.Run() -- means
		// that fresh database tables and data will be set up and torn down
		// for each sub-test.
		db := newTestDB(t)
		// Create a new instance of the UserModel.
		m := TaskModel{db}
		// Call the UserModel.Exists() method and check that the return
		// value and error match the expected values for the sub-test.
		tasks, err := m.GetAll()

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

		assert.Equal(t, &tasks, &expectedTasks)
		if tasks == nil {
			assert.ErrNoRows(t, err)
		}
	})
}
