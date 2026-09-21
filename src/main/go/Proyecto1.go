package main

import(
	"os"
)

func main(){
	aplicacion := NewAplicacion(os.Args[1:])
	aplicacion.Ejectuar()
}
