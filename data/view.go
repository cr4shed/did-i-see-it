package data

import (
	"fmt"
	"database/sql"
)

func GetViewsByCollection(db *sql.DB, collectionId int) ([]View, error) {
    rows, err := db.Query("SELECT Id, CollectionId, MediaId FROM View WHERE CollectionId = ?", collectionId)
	if err != nil {
		return nil, fmt.Errorf("ERROR - Could not query database. %v", err)
	}

	defer rows.Close()
	
    var views []View
	for rows.Next() {
		var view View

        err := rows.Scan(&view.Id, &view.CollectionId, &view.MediaId)
        if err != nil {
            return nil, fmt.Errorf("ERROR - Could not scan View row. %v", err)
        }

        views = append(views, view)
	}

    return views, nil
}