package controllers

import (
	"alura-go-restapi/database"
	"alura-go-restapi/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func BuscarPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	err := json.NewEncoder(w).Encode(personalidades)
	if err != nil {
		panic(err.Error())
	}
}

func BuscarPersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	err := json.NewEncoder(w).Encode(personalidade)
	if err != nil {
		panic(err.Error())
	}
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	if err != nil {
		panic(err.Error())
	}
	database.DB.Create(&novaPersonalidade)
	err = json.NewEncoder(w).Encode(&novaPersonalidade)
	if err != nil {
		panic(err.Error())
	}
}

func ExcluirPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	resultado := database.DB.First(&personalidade, id)

	if resultado.RowsAffected == 0 || resultado.Error != nil {
		err := json.NewEncoder(w).Encode("Personalidade não encontrada")
		if err != nil {
			panic(err.Error())
		}
		panic(resultado.Error)
	}
	database.DB.Delete(&personalidade, id)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	resultado := database.DB.First(&personalidade, id)

	if resultado.RowsAffected == 0 || resultado.Error != nil {
		err := json.NewEncoder(w).Encode("Personalidade não encontrada")
		if err != nil {
			panic(err.Error())
		}
		panic(resultado.Error)
	}

	err := json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)

	err = json.NewEncoder(w).Encode(&personalidade)
	if err != nil {
		panic(err.Error())
	}
}
