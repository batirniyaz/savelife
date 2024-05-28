// db.go

package database

import (
    "database/sql"
    "log"
    
    _ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// DB is the global database connection
var DB *sql.DB

// Init initializes the database connection
func Init(dbPath string) error {
    var err error
    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        return err
    }
    if err := DB.Ping(); err != nil {
        return err
    }
    log.Println("Connected to database")
    return nil
}

// Close closes the database connection
func Close() {
    if DB != nil {
        DB.Close()
    }
}
