package main

import(
	"os"
)

func main(){
	aplicacion := newAplicacion(os.Args[1:])
	aplicacion.ejecutar()
}
