package main

import (
	"fmt"
	"os"
	"strconv"
)

// go build -o proyecto1
// go run main.go
func uso(){
	fmt.Fprintln(os.Stderr, "Uso: ./proyecto1" + "puerto")
	os.Exit(1)
}

func errorPuerto(p string){
	fmt.Fprintln(os.Stderr, "El puerto %s es inválido", p)
	os.Exit(1)
}

type Aplicacion struct{
	args []string
	puerto string
}

func NewAplicacion(args []string) *Aplicacion{
	puerto := "5050"

	if len(args) != 1{
		uso()
	} 
	if (!verificaPuerto(args[0])){
		errorPuerto(args[0])
	}

	puerto = args[0]
	return &Aplicacion{
		args: args,
		puerto: puerto,
	}
}

func (app *Aplicacion) Ejecutar() {
	direccion := ":" + app.puerto
	server := NewServidor(direccion)
	server.iniciaServidor()
}


func (app *Aplicacion) validaPuerto(puerto string) bool{
	p, err := strconv.Atoi(puerto)
	if err != nil{
		return false
	}
	if (p < 1025 || p > 35535){
		return false
	}
	return true
}