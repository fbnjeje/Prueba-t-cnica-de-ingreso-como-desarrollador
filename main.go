package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type video struct {
	ID            int    `json:ID`
	Tema          string `json:Tema`
	fechaCreacion string `json:fechaCreacion`
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

func createVideo(w http.ResponseWriter, r *http.Request) {
	var newVideo video

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "inserte un video valido")
	}

	json.Unmarshal(reqBody, &newVideo)

	append(videos, newVideo)
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(videos)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}

func main() {

	// router := mux.NewRouter().StrictSlash(true
	router := mux.NewRouter().StrictSlash(true)

	//Rutas

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/videos", getVideos)

	log.Fatal(http.ListenAndServe(":3000", router))
}
