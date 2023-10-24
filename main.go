package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type video struct {
	ID            int    `json:ID`
	Tema          string `json:Tema`
	fechaCreacion string `json:Creacion`
	Creador       string `json:Creador`
	Descripcion   string `json:Descripcion`
	Titulo        string `json:Titulo`
	Estado        bool   `json:Estado`
}

type allVideos []video

var videos = allVideos{
	{
		ID:            1,
		Tema:          "Matematcias",
		fechaCreacion: "23/10/2023",
		Creador:       "JulioProfe",
		Descripcion:   "Aprenderas a hacer una calculadora",
		Titulo:        "Aprende a como solucionar una calculadora",
		Estado:        true,
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
}
