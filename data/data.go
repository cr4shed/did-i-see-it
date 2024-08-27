package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func DbConnect() (*sql.DB, error) {
    // Connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    os.Getenv("DBNET"),
        Addr:   os.Getenv("DBADDR"),
        DBName: os.Getenv("DBNAME"),
        AllowNativePasswords: true,
    }

    // Check connection.
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, fmt.Errorf("ERROR - Could not connect to database. %v", err)
    }

    // Ping database.
    pingErr := db.Ping()
    if pingErr != nil {
        return nil, fmt.Errorf("ERROR - Failed to ping database. %v", pingErr)
    }

    fmt.Println("SUCCESS - Connected to database.")

    return db, nil
}