:
import (
       "net"   
       "log"
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
     return &server
}

// ESTO FUE TOMADO DE GOBYEXAMPLE/TCP SERVER.COM  
func iniciaServidor(){
     server.conexionActiva = true
     //como llamar a variables de clase
     enviaMensajePublico("El puerto %d está disponible", puerto) 

     server.serverSocket, err := net.Listen("tcp", puerto)
     
     if err != nil{ //// por que dice que tiene que ser distinto. REVISAR
          log.Fatal("Error al escuchar:", err) /// debe ser de tipo mensaaje.  CAMBIAR
     }
     defer listener.Close()


     for{
          conn, err := serverSocket.Accept()
          if err != nil{
               log.Println("Error al aceptar la conexion", err)
          }

          go RecibeMensajes(conn, //revisar si tambien debe recibir algo mas)

     }
}

func agregaConexion(){
     ## crear Sala principal

}

func terminaConexion(conn conexion){

}


func enviaMensajePublico(){

}


func enviaMensajeUsuario(){

}


func enviaMensajeSala(){

}



func RecibeMensajes(conn conexion, mensaje string){

     // mando a llamara recibePeticion() y recibeMensaje segun el caso
}
func recibePeticion(){

}


func recibeMensaje(conn conexion, mensaje string){

}


func verificaUsuario(nombreUsuario string){

}




func creaSala(nombreSala string){

}

func salasActivas(){

}

func usuariosActivos(){

}




