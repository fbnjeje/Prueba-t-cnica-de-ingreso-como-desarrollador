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
	ID          int    `json:ID`
	Tema        string `json:Tema`
	Creador     string `json:Creador`
	Descripcion string `json:Descripcion`
	Titulo      string `json:Titulo`
	Estado      bool   `json:Estado`
	CreatedOn   string `json:CreatedOn`
}

type allVideos []video

var videos = allVideos{
	{
		ID:          1,
		Tema:        "Matematcias",
		Creador:     "JulioProfe",
		Descripcion: "Aprenderas a hacer una calculadora",
		Titulo:      "Aprende a como solucionar una calculadora",
		Estado:      true,
		CreatedOn:   "24/10/2023",
	},
}

func createVideo(w http.ResponseWriter, r *http.Request) {
	var newVideo video

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "inserte un video valido")
	}

	json.Unmarshal(reqBody, &newVideo)

	newVideo.ID = len(videos) + 1

	videos = append(videos, newVideo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newVideo)
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	router.HandleFunc("/videos", createVideo)

	log.Fatal(http.ListenAndServe(":3000", router))
}
