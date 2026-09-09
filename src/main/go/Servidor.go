
import (
       "net"   
)


type servidor struct{
     puerto int
     conexiones map[int]conexion
     salas [string]Sala
     serverSocket net.Listener 
     conexionActiva bool
}

func NewServidor(puerto int) *servidor {
     server := servidor{puerto: puerto}
     server.conexiones = make(map[int]conexion)
     server.salas = make[map[string]Sala]
     server.serverSocket = net.Listen("tcp", puerto)
     return &server
}


func iniciaServidor(){
     server.conexionActiva = true 
}

func agregaConexion(){
     ## crear Sala principal

}

func verificaUsuario(nombreUsuario string){

}

func creaSala(nombreSala string){

}

func salasActivas(){

}

func usuariosActivos(){

}

func terminaConexion(){

}



