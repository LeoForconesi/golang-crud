package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoforconesi/restApiCRUD/db"
	"github.com/leoforconesi/restApiCRUD/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	log.Println("Buscando tasks")
	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
	w.WriteHeader(http.StatusOK)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task // 1ro crear el elemento para hacer referencia y mapear el body del request
	// var user models.User
	// en el body va a venir un user id al cual crearle la tarea, hay que verificar si existe o no antes de crear la tarea.
	json.NewDecoder(r.Body).Decode(&task) // decodear el body del request al tipo task

	taskCreated := db.DB.Create(&task) // creo el elemento en la base de datos, gorm ya sabe de que tabla estamos hablando al pasarle el puntero de task

	if taskCreated.Error != nil { // verifico si hubo algun error al crear
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(taskCreated.Error.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)     // respondo con el status ok
	json.NewEncoder(w).Encode(&task) // encodeo a JSON la respuesta del elemento creado
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	params := mux.Vars(r)

	db.DB.First(&task, params["id"]) // es el id de la tarea
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tassk not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // el sxtatus no content es no tengo nada que devolver pero la tarea fue bien
}
