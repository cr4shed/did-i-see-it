package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cr4shed/did-i-see-it/data"
	"github.com/joho/godotenv"
)

func main() {
	envErr := handleLoadEnv("../.env")
    if envErr != nil {
        log.Fatal(envErr)
    }

	db, dbErr := data.DbConnect()
    if dbErr != nil {
        log.Fatal(dbErr)
    }

	test(db)
}

func handleLoadEnv(envPath string) (error) {
    // Load env variables from .env file.
    err := godotenv.Load(envPath)
    if err != nil {
        return fmt.Errorf("ERROR - Could not load .env file. %v", err)
    }

    return nil
}

func test(db *sql.DB) {
	_, err := data.AddCollection(db, 1, data.Collection{Name: "Test Collection From the DAL"})
	if err != nil {
		log.Fatal(err)
	}

	collection1, errc1 := data.GetUserCollections(db, 1)
	if errc1 != nil {
		log.Fatal(errc1)
	}

	collection2, errc2 := data.GetUserCollections(db, 2)
	if errc2 != nil {
		log.Fatal(errc2)
	}

	views1, err1 := data.GetViewsByCollection(db, 1)
	if err1 != nil {
		log.Fatal(err1)
	}
	views2, err2 := data.GetViewsByCollection(db, 2)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(views1)
	fmt.Println(views2)
	fmt.Println(collection1)
	fmt.Println(collection2)
}