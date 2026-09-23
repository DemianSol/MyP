package main

import (
	"fmt"
	"os"
	"strconv"
)

func validaPuerto(puerto string) bool{
	p, err := strconv.Atoi(puerto)
	if err != nil{
		return false
	}
	if (p < 1025 || p > 65535){
		return false
	}
	return true
}

// go build -o proyecto1
// go run main.go
func uso(){
	fmt.Fprint(os.Stderr, "Uso: ./proyecto1 <puerto>\n")
	os.Exit(1)
}

func errorPuerto(p string){
	fmt.Fprintf(os.Stderr, "El puerto %s es inválido\n", p)
	os.Exit(1)
}

type aplicacion struct{
	args []string
	puerto string
}

func newAplicacion(args []string) *aplicacion{
	puerto := "5050"

	if len(args) != 1{
		uso()
	} 
	if (! validaPuerto(args[0])){
		errorPuerto(args[0])
	}

	puerto = args[0]
	return &aplicacion{
		args: args,
		puerto: puerto,
	}
}

func (app *aplicacion) ejecutar() {
	direccion := ":" + app.puerto
	server := newServidor(direccion)
	server.iniciaServidor()
}


