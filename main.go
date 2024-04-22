package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoforconesi/restApiCRUD/db"
	"github.com/leoforconesi/restApiCRUD/models"
	"github.com/leoforconesi/restApiCRUD/routes"
)

func main() {
	fmt.Println("Start")

	db.DBConnection()

	// Migracion de los modelos de tabla usando gorm. Indico los structs que tiene que migrar.
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	// Task Routes

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
