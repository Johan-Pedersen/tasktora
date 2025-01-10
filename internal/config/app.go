package config

import (
	"log"
	"tasktora/internal/models"
)

/*
Information which should be available for handlers
Like logging and DB connection pools
*/
type App struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	TaskModel   models.ITaskModel
}
