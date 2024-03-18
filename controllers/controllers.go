package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/helioLJ/golang-rest-api-personalities/models"
	"github.com/helioLJ/golang-rest-api-personalities/database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, seja bem-vindo!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}


func CriaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.Delete(&models.Personalidade{}, id)
	fmt.Fprintf(w, "A personalidade com o ID %v foi deletada com sucesso", id)
}