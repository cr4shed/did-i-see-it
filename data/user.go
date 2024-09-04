package data

import (
	"database/sql"
	"fmt"
	"os"

    "golang.org/x/crypto/bcrypt"
	//"github.com/golang-jwt/jwt/v5"
)

var salt string

func CreateUser(db *sql.DB, username string, email string, password string) (int, error) {
	usernameExists, usernameErr := CheckUsernameExists(db, username)
	if usernameErr != nil {
		return BAD_INT, fmt.Errorf("could not check if username exists. %v", usernameErr)
	}

	emailExists, emailErr := CheckEmailExists(db, email)
	if emailErr != nil {
		return BAD_INT, fmt.Errorf("could not check if email exists. %v", emailErr)
	}
	
	if usernameExists || emailExists {
		return BAD_INT, fmt.Errorf("username or email already exists")
	}

	id, insertErr := insertUser(db, username, email, password)
	if insertErr != nil {
		return BAD_INT, fmt.Errorf("could not insert user. %v", insertErr)
	}

	return int(id), nil
}

func LoginUser(db *sql.DB, username string, password string) error {
	var hash string

	err := db.QueryRow("SELECT PassHash FROM User WHERE Username = ?", username).Scan(&hash)
	if err != nil {
		return fmt.Errorf("could not query User table. %v", err)
	}

	if !checkPasswordMatch(hash, password) {
		return fmt.Errorf("password does not match")
	}

	return nil
}

func getHashPassword(password string) (string, error) {
	if salt == "" {
		salt = os.Getenv("SALT")
		
		if salt == "" {
			return "", fmt.Errorf("salt is not set")
		}
	}

    bytePassword := []byte(password + salt)

	// Ensure password does not exceed the byte limit for bcrypt.
	if len(bytePassword) > 72 {
		return "", fmt.Errorf("password is too long")
	}

    hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    return string(hash), nil
}

func CheckUsernameExists(db *sql.DB, username string) (bool, error) {
	var exists bool

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE Username = ?)", username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not query User table. %v", err)
	}

	return exists, nil
}

func CheckEmailExists(db *sql.DB, email string) (bool, error) {
	var exists bool

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE Email = ?)", email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not query User table. %v", err)
	}

	return exists, nil
}

func insertUser(db *sql.DB, username string, email string, password string) (int, error) {
	hash, err := getHashPassword(password)
	if err != nil {
		return BAD_INT, err
	}

	result, err := db.Exec("INSERT INTO User (Username, Email, PassHash) VALUES (?, ?, ?)", username, email, hash)
	if err != nil {
		return BAD_INT, fmt.Errorf("could not insert into User. %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return BAD_INT, fmt.Errorf("record inserted into User but could not get last insert id. %v", err)
	}

	return int(id), nil
}

func checkPasswordMatch(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)

	return err == nil
}