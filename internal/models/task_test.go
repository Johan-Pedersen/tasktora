package models

import (
	"database/sql"
	"fmt"
	"log"
	"tasktora/internal/assert"
	"testing"
	"time"
)

func TestTaskModel(t *testing.T) {
	time, err1 := time.Parse("2006-01-02 15:04:05", "2025-01-01 10:00:00")
	if err1 != nil {
		log.Fatal("dooo")
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
				Created:  time,
				ParentId: sql.NullInt64{Valid: false},
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
				Created:  time,
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
			fmt.Printf("task: %v\n", task)
			assert.Equal(t, task, tt.want)
			if task == nil {
				assert.ErrNoRows(t, err)
			}
		})
	}
}
