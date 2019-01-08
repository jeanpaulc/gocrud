package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"time"
	"os"

	// ROUTING
	"github.com/gorilla/mux"

	// LOGGING
	"github.com/gorilla/handlers"

	// PSQL DATABASE ORM
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User - This will be the model for a user
type User struct {
	ID        uint       	`gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	FirstName string     	`json:"first_name"`
	LastName  string     	`json:"last_name"`
	Email     string     	`json:"email"`
	Age       json.Number	`json:"age,Number"`
	CreatedAt time.Time  	`json:"created_at"`
	UpdatedAt time.Time  	`json:"updated_at"`
}

var db *gorm.DB

func main() {
	// openDB opens the Postgres DB and Auto Migrates the latest table schema
	openDB()
	defer db.Close()

	router := mux.NewRouter()
	router.Use(setHeader)

	// CRUD ROUTES
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Logs requests to out in Apache Common Log Format 
	logger := handlers.LoggingHandler(os.Stdout, router)

	log.Fatal(http.ListenAndServe(":3000", logger))
}

func openDB() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost dbname=gocrud sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema; Only makes changes to the table rows not the table data
	// NOTE: for dev purposese this is run every time the router is started
	db.AutoMigrate(&User{})
}

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

// getUsers - Get all users from db
func getUsers(res http.ResponseWriter, req *http.Request) {
	// set users to an empty slice (array)
	var users []User
	//GORM's db.Find() creates the SQL query `SELECT * FROM users;`
	db.Find(&users)
	// Encode into JSON and respond with all users
	json.NewEncoder(res).Encode(&users)
}

// getUser - Get single user from db
func getUser(res http.ResponseWriter, req *http.Request) {
	// Pull the id from the request url
	id := mux.Vars(req)["id"]

	var user User
	//GORM's db.First() creates the SQL query `SELECT * FROM users WHERE id = $1`
	db.First(&user, id)
	// Encode into JSON and respond with user query
	json.NewEncoder(res).Encode(&user)
}

// createUser - create a single user in db
func createUser(res http.ResponseWriter, req *http.Request) {
	var data User
	json.NewDecoder(req.Body).Decode(&data)

	fName := data.FirstName
	lName := data.LastName
	email := data.Email
	age := data.Age

	newUser := &User{FirstName: fName, LastName: lName, Email: email, Age: age}

	db.Create(newUser)
	// Encode into JSON and respond with newly created user
	json.NewEncoder(res).Encode(newUser)
}

func updateUser(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	var data User
	json.NewDecoder(req.Body).Decode(&data)

	fName := data.FirstName
	lName := data.LastName
	email := data.Email
	age := data.Age

	updatedUser := User{}

	for 

	var user User
	//GORM's db.First() creates the SQL query `SELECT * FROM users WHERE id = $1`
	db.First(&user, id)
	// Update multiple attributes with `struct`, will only update those changed & non blank fields
	db.Model(&user).Updates(updatedUser)
	// Encode into JSON and respond with updated user
	json.NewEncoder(res).Encode(updatedUser)
}

// deleteUser - create a single user in db
func deleteUser(res http.ResponseWriter, req *http.Request) {
	// Pull the id from the request url
	id := mux.Vars(req)["id"]

	var user User
	//GORM's db.First() creates the SQL query `SELECT * FROM users WHERE id = $1`
	db.First(&user, id)
	db.Delete(&user)
	// Encode into JSON and respond with deleted user
	json.NewEncoder(res).Encode(&user)
}
