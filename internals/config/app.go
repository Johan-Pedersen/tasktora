package config

import (
	"log"
)

/*
Information which should be available for handlers
Like logging and DB connection pools
*/
type App struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}
