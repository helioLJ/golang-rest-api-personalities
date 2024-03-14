package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/helioLJ/golang-rest-api-personalities/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, seja bem-vindo!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
