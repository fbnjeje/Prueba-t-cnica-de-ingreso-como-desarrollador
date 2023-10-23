package main

import (
	"fmt"
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

func main() {
	fmt.Print("hola como estan")
}
