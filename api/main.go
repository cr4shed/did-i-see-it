package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/cr4shed/did-i-see-it/data"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load env variables.
	envErr := handleLoadEnv("../.env")
	if envErr != nil {
		log.Fatal(envErr)
	}
}

func main() {
	// Connect to the database.
	db, dbErr := data.DbConnect()
    if dbErr != nil {
        log.Fatal(dbErr)
    }

	// Setup API routing.
	router := gin.Default()
	router.Use(databaseMiddleware((db)))

	router.GET("/collections/:userId", getUserCollections)

	router.POST("/auth/register", registerUser)


	router.Run("localhost:8080")
}

func handleLoadEnv(envPath string) (error) {
    // Load env variables from .env file.
    err := godotenv.Load(envPath)
    if err != nil {
        return fmt.Errorf("ERROR - Could not load .env file. %v", err)
    }

    return nil
}

func databaseMiddleware(db *sql.DB) gin.HandlerFunc {
	err := db.Ping()

	// If database connection is in a bad state, abort.
	if err != nil {
		return func (c *gin.Context) {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
		}
	}

	// Continue processing.
	return func (c *gin.Context) {
		c.Set(DATABASE_CONNECTION, db)
		c.Next()
	}
}

func handleResponse[K, V Returnable](c *gin.Context, obj V, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.JSON(http.StatusOK, obj)
	}
}

func getUserCollections(c *gin.Context) {
	db, _ := c.Get(DATABASE_CONNECTION)

	userId := c.Param("userId")
	collection, err := data.GetUserCollections(db.(*sql.DB), userId)

	handleResponse[[]data.Collection](c, collection, err)
}

func registerUser(c *gin.Context) {
	db, _ := c.Get(DATABASE_CONNECTION)

	var user UserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := data.CreateUser(db.(*sql.DB), user.Username, user.Email, user.Password)

	handleResponse[IdResposne](c, IdResposne{Id: id}, err)
}