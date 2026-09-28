package main

import(
	"os"
)

// https://gobyexample.com/command-line-arguments
func main(){
	aplicacion := newAplicacion(os.Args[1:])
	aplicacion.ejecutar()
}
