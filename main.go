package main

import (
	"fmt"

	"github.com/helioLJ/golang-rest-api-personalities/models"
	"github.com/helioLJ/golang-rest-api-personalities/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
