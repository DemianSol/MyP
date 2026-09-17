// modificación estructura/diseño   --- > agregar al siguiente commit 
// cambio en inicia servidor      ----- > agregar al sigueinte commit 
import (
       "net"   
       "log"
       "fmt"
       "bufio"
       "os"
)


type servidor struct{
     puerto string
     enchufe net.Listener
     conexiones map[int]Conexion
     clientes map[string]Cliente // cambiar para manejar clase de Vala
     salas map[string]sala
     conexionActiva bool
     buzonTareas chan tarea 
}

func NewServidor(puerto int *servidor) {
     server := servidor{puerto: puerto}
     server.conexiones = make(map[int]conexion)
     server.salas = make(map[string]Sala)
     server.cliente = make(map[int]Cliente)
     server.lectura = bufio.NewScanner(os.Stdin)
     buzonTareas = make(chan <-tarea, 1000) //buscar implementacion dinamica
     return &server
}

// ESTO FUE TOMADO DE GOBYEXAMPLE/TCP SERVER.COM  
func (s *servidor) iniciaServidor() error{
     var err error

     go s.manejaBuzon()

     s.conexionActiva = true
     imprimeMensaje("El puerto %s está disponible", s.puerto) 

     while(conexionActiva){
          s.enchufe, err = net.Listen("tcp", s.puerto)   
          if err != nil{ 
               return fmt.Errorf("Error al escuchar el puerto %d: %w", s.puerto, err) 
          }
          defer listener.Close()
          imprimeMensaje("Servidor escuchando")    
          conn, err := enchufe.Accept()
          if err != nil{
               fmt.Errorf("Error al aceptar la conexion", err) 
               continue
          }
          conexion := conexion{conn, buzonTareas}
          imprimeMensaje("Conexion recibida de") // agregar información 
          
          go (conexion.recibeMensaje()) // puedo hacerlo mas especifico para asegurar que recibaMensajeCliente
               //sincronizar          
     }
}


func imprimeMensaje (mensaje string){ //probablemente está mal
     fmt.Println(mensaje)
}

func (s *servidor) manejaBuzon() error{ //errores en hilos
     var err error
     for tarea := range s.buzonTareas{
          s.procesaMensaje(tarea.getConexion(), tarea.getMensaje())
     }
     
     if err != nil{
          fmt.Errof("Error al procesar el mensaje: %w", err)
     }
}

func (s servidor) procesaMensaje(conexion Conexion, mensaje Mensaje) error{
     if (! conexion.getConexionActiva())
          return
     switch mensaje {

     case IDENTIFY:
          salasActivas
     case STATUS: 

     case USERS:

     case TEXT:

     case PUBLIC_TEXT:

     case NEW_ROOM:


     case INVITE:

     case JOIN_ROOM:

     case ROOM_USERS:

     case ROOM_TEXT:

     case LEAVE_ROOM:

     case DISCONNECT:
          
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





func verificaUsuario(nombreUsuario string){

}




func creaSala(nombreSala string){

}

func salasActivas(){

}

func usuariosActivos(){

}




