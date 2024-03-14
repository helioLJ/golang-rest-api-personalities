package routes

import (
	"log"
	"net/http"

	"github.com/helioLJ/golang-rest-api-personalities/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
