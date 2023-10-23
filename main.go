package main

import (
	"fmt"
)

type video struct {
	ID int `json:ID`
	Tema
	fechaCreacion
	Creador
	Descripcion
	Titulo
	Estado
}

func main() {
	fmt.Print("hola como estan")
}
