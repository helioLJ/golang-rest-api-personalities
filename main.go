package main

import (
	"fmt"

	"github.com/helioLJ/golang-rest-api-personalities/database"
	"github.com/helioLJ/golang-rest-api-personalities/models"
	"github.com/helioLJ/golang-rest-api-personalities/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Id: 2, Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
