package data

import (
	"database/sql"
	"fmt"
)

type Collection struct {
	Id		  		int
	Name			string
}

func GetUserCollections(db *sql.DB, userId string) ([]Collection, error) {
    rows, err := db.Query("SELECT Id, Name FROM Collection WHERE UserId = ?", userId)
    if err != nil {
        return nil, fmt.Errorf("ERROR - Could not query database. %v", err)
    }

    defer rows.Close()

    var collections []Collection
    for rows.Next() {
        var collection Collection

        err := rows.Scan(&collection.Id, &collection.Name)
        if err != nil {
            return nil, fmt.Errorf("ERROR - Could not scan Collection row. %v", err)
        }

        collections = append(collections, collection)
    }

    return collections, nil
}

func AddCollection(db *sql.DB, userId int, collection Collection) (int64, error) {
	result, err := db.Exec("INSERT INTO Collection (UserId, Name) VALUES (?, ?)", userId, collection.Name)
	if err != nil {
		return -1, fmt.Errorf("ERROR - Could not insert into Collection. %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("ERROR - Record inserted into Collection but could not get last insert id. %v", err)
	}

	return id, err;
}