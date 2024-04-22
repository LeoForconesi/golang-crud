package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoforconesi/restApiCRUD/db"
	"github.com/leoforconesi/restApiCRUD/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
	// w.Write([]byte("get usersss"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r) // params es un map, puedo acceder a los datos del mapa accediendo su key, en este caso params["id"]
	fmt.Println("Buscando al userID: " + params["id"])

	db.DB.First(&user, params["id"])

	// si el usuario no existe, go por defecto devuelve 0, si buscas un string devuelve vacio, es la forma de go de avisar que lo pedido no lo encuentra
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// con esto hago la asociacion, donde traigo la task que coincida con el userid. Tanto task como user tienen userid.
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
	// w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// mapea el body del request a un objeto tipo user
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("Creando user")

	// guarda el user en la base de datos
	createdUser := db.DB.Create(&user)

	// verifica que no haya errores, si hubiera alguno, devuelve un error 400 en este caso
	err := createdUser.Error
	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		w.Write([]byte(err.Error()))
	}

	// serializa la respuesta a JSON.
	json.NewEncoder(w).Encode(&user)
	// w.Write([]byte("post user"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not Found"))
		return
	}
	log.Println("Eliminando al userID: ", user.ID)
	db.DB.Delete(&user) // este delete hace un delete logico
	// db.DB.Unscoped().Delete(&user) // este hace un delete total
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Delete user"))
}
