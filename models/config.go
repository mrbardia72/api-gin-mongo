package models

import (
	"os"

	"github.com/mrbardia72/api-gin-mongo/db"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE")

// Database name
var DatabaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var DBConnect = db.NewConnection(server, DatabaseName)
