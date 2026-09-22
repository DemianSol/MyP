// modificación estructura/diseño   --- > agregar al siguiente commit 
// cambio en inicia servidor      ----- > agregar al sigueinte commit 
import (
       "net"   
       "log"
       "fmt"
       "bufio"
       "os"
)
type Status string

const(
     AWAY Status : "AWAY"
     ACTIVE Status: "ACTIVE"
     BUSY Status: "BUSY"
)

type servidor struct{
     puerto string
     enchufe net.Listener
     contadorConexiones int
     conexiones map[*Conexion]*Usuario
     clientes map[string]*Usuario // cambiar para manejar clase de Vala
     salas map[string]*Sala
     conexionActiva bool
     buzonTareas chan tarea 
}

func NewServidor(puerto int) *servidor{
     server := servidor{puerto: puerto}
     server.contadorConexiones = 0
     server.conexiones = make(map[*Conexion]*Usuario)
     server.salas = make(map[string]Sala)
     server.clientes = make(map[string]*Usuario)
     server.lectura = bufio.NewScanner(os.Stdin)
     server.buzonTareas = make(chan <-tarea, 1000) 
     return &server
}

// ESTO FUE TOMADO DE GOBYEXAMPLE/TCP SERVER.COM  
func (s *servidor) iniciaServidor() error{  
     s.conexionActiva = true
     s.imprimeMensaje("El puerto %s está disponible", s.puerto) 

     s.enchufe, err = net.Listen("tcp", s.puerto)   
     if err != nil{ 
               return fmt.Errorf("Error al escuchar el puerto %s: %w", s.puerto, err) 
          }
     defer s.enchufe.Close()
     go s.manejaBuzon()
     
     s.imprimeMensaje("Servidor escuchando") // creo que está mal   

     for s.conexionActiva{
          conn, err := enchufe.Accept()
          if err != nil{
               continue
          }
          s.contadorConexiones++
          conexion := NewConexion(conn, s.contadorConexiones)
     
          s.imprimeMensaje("Conexion recibida de: %d", conexion.getIdentificador()) // agregar información 
          
          go (s.recibeMensajes(conexion)) 
     }
     return nil
}


func imprimeMensaje(mensaje string){ //probablemente está mal
     fmt.Println(mensaje)
}

func (s *servidor) recibeMensajes(con *Conexion){
     // que hago cuando se termine? debo mandar mensaje de desconexion? 
     for {
          bytes, err := con.recibeMensaje()
          if err != nil{
               break
          }
          msj, err := ProcesaJSON(bytes)
          if err != nil{
               continue
          }
          s.buzonTareas <- tarea{conn: con, mensaje: msj}
     }
}

// https://blog.joedayz.pe/channels-en-go-buffered-vs-unbuffered-comunicacion-segura-entre-goroutines
func (s *servidor) manejaBuzon() error{ //errores en hilos
     var err error
     for tarea := range s.buzonTareas{
          s.procesaMensaje(tarea.getConexion(), tarea.getMensaje())
     }
     
     if err != nil{
          return fmt.Errof("Error al procesar la petición: %w", err)
     }
}
// https://go.dev/doc/effective_go#type_switch
func (s *servidor) procesaMensaje(conexion *Conexion, mensaje Mensaje) error{
     if (! conexion.getConexionActiva())
          return nil

     switch msj := mensaje.(type) {

     case mensajeIdentificar:

          _, user := s.conexiones[conexion]
          if (user){
               return nil // que debo hacer si un usuario ya identificado intenta volver a identificarse?
          }

          if (s.verificaUsuario(msj.Username)){
               resp := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("IDENTIFY", "USER_ALREADY_EXISTS", msj.Username)
               s.enviaMensajeUsuario(conexion, resp)
               return nil
          }

          nuevoUsuario := NewUsuario(msj.Username, conexion)
          s.clientes[msj.Username] = nuevoUsuario
          s.conexiones[conexion] = nuevoUsuario
          conexion.setAceptado(true)


          res := NewMensajeAClienteBuilder(respuesta).
          añadirRespuesta("IDENTIFY", "SUCCESS", msj.Username)
          s.enviaMensajeUsuario(conexion, res)

          todos := NewMensajeAClienteBuilder(nuevoUsuario).
          añadir("username", msj.Username)
          s.enviaMensajePublico(todos, conexion)
               
          

     case mensajeNewStatus: 
          user, b := s.conexiones[conexion]
          if !b{
               conexion.desconectar()
               return nil 
          }

          estado := Status(msj.Status)
          if (estado != ACTIVE && estado != AWAY && estado != BUSY){
               return nil
          }
          if (user.getStatus() != estado){
               user.setStatus(estado)
               builder := NewMensajeAClienteBuilder(nuevoStatus).
               añadir("username", user.getNombre()).
               añadir("status", string(estado))
               s.enviaMensajePublico(builder, conexion)
          }


     case mensajeListaUsuarios: 
          _, registrado := s.conexiones[conexion]
          if !registrado {
            conexion.Desconectar()
            return nil
          }
          u := make(map[string]string)
          for _, user := range s.clientes{
               u[user.getNombre()] = string(user.getStatus())
          }
          builder := NewMensajeAClienteBuilder(listaUsuario).
          añadeListaUsuarios(u)
          s.enviaMensajeUsuario(conexion, builder)

     case mensajeTexto:
          emisor, b := s.conexiones[conexion]
          if !b{
               conexion.desconectar()
               return nil 
          }

          destinatario, user := s.clientes[msj.Username]
          if !user {
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("TEXT", "NO_SUCH_USER", msj.Username)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }


          destino := NewMensajeAClienteBuilder(textoDesde).
          añadir("username", emisor.getNombre()).
          añadir("text", msj.Text)
          s.enviaMensajeUsuario(destinatario.getConexion(), destino)
          }



     case mensajeTextoPublico:
          remitente, b := s.conexiones[conexion]
          if !b {
               conexion.desconectar()
               return nil 
          }

          msg := NewMensajeAClienteBuilder(textoPublico).
          añadir("username", remitente.getNombre()).
          añadir("text", msj.Text)

          s.enviaMensajePublico(msg, remitente.getConexion())

     case mensajeNuevaSala:
          
          sala, existe := s.salas[msj.Roomname]

          if !existe {
               remitente, b := s.conexiones[conexion]
               if !b{
                    coenxion.desconectar()
               }
               room := NewSala(msj.Roomname, remitente)
               s.salas[msj.Roomname] = room
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("NEW_ROOM", "SUCCESS", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }

          answer := NewMensajeAClienteBuilder(respuesta).
          añadirRespuesta("NEW_ROOM", "ROOM_ALREADY_EXISTS", msj.Roomname)
          s.enviaMensajeUsuario(conexion, answer)
          return nil
          
     case mensajeInvita:
          remitente, b := s.conexiones[conexion]
          if !b {
               conexion.Desconectar()
               return nil 
          }


          sala, existe := s.salas[msj.Roomname]

          if(! existe){
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("INVITE", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }
          
          if (! enSala(remitente.getNombre())){
               return nil
           }
          
          for _, nombre := range msj.Users{
               _, answer := s.clientes[nombre]
               if (!answer){
                    respuesta := NewMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("INVITE", "NO_SUCH_USER", user.getNombre())
                    s.enviaMensajeUsuario(conexion, respuesta)
                    return nil
               }
          }

          invitacion := NewMensajeAClienteBuilder(invitacion).
          añadir("username", remitente.getNombre()).
          añadir("roomname", msj.Roomname)

          for _, user := range msj.Users{
               if (sala.enSala(user) || sala.estaInvitado(user)){
                    continue
               }
               sala.invitacion(user)
               clientConn, _ := s.clientes[user]
               s.enviaMensajeUsuario(clientConn.getConexion(), invitacion)
          }
          return nil

          
     case mensajeUnirSala:
          remitente, b := s.conexiones[conexion]
          if !b{
               conexion.desconectar()
               return nil 
          }
          
          sala, existe := s.salas[msj.Roomname]

          if (!existe){
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("JOIN_ROOM", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }

          if (!sala.estaInvitado(remitente.getNombre())){
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("JOIN_ROOM", "NOT_INVITED", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }

          respuesta := NewMensajeAClienteBuilder(respuesta).
          añadirRespuesta("JOIN_ROOM", "SUCCESS", msj.Roomname)
          s.enviaMensajeUsuario(conexion, respuesta)

          sala.agregarCliente(remitente)

          msjPublico := NewMensajeAClienteBuilder(unidoASala).
          añadir("roomname", msj.Roomname).
          añadir("username", remitente.getNombre())

          sala.mensajePublico(msjPublico)
          return nil

     case mensajeUsuariosSala:
          remitente, b := s.conexiones[conexion]
          if !b{
               conexion.desconectar()
               return nil 
          }
          sala, existe := s.salas[msj.Roomname]

          if (!existe){
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("ROOM_USRS", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }

          if(!sala.enSala(remitente.getNombre())){
               respuesta := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("ROOM_USERS", "NOT_JOINED", msj.Roomname)
               s.enviaMensajeUsuario(conexion, respuesta)
               return nil
          }

          lista := make(map[string]string)
          for _, u := range sala.getUsuarios() {
            lista[u.getNombre()] = string(u.getStatus())
        }
          respuesta := NewMensajeAClienteBuilder(usuariosSala).
          añadir("roomname", msj.Roomname).
          añadeListaUsuarios(lista)
          s.enviaMensajeUsuario(conexion, respuesta)
          return nil

     case mensajeTextoSala:
          emisor, registrado := s.conexiones[conexion]
          if !registrado {
               conexion.desconectar()
               return nil 
          }

          sala, existe := s.salas[msj.Roomname]
          if !existe {
               resp := NewMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("ROOM_TEXT", "NO_SUCH_ROOM", msj.Roomname)
                    s.enviaMensajeUsuario(conexion, resp)
                    return nil
          }

          if (! sala.enSala(emisor.getNombre())){     
               resp := NewMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("ROOM_TEXT", "NOT_JOINED", msj.Roomname)
                    s.enviaMensajeUsuario(conexion, resp)
                    return nil
          }
          notif := NewMensajeAClienteBuilder(textoDesdeSala).
               añadir("roomname", msj.Roomname).
               añadir("username", emisor.nombre).
               añadir("text", msj.Text)

          sala.mensajePublico(notif, emisor)

     case mensajeDejaSala:
          emisor, registrado := s.conexiones[conexion]
          if !registrado {
               conexion.desconectar()
               return nil 
          }

          sala, existe := s.salas[msj.Roomname]
          if !existe {
            resp := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("LEAVE_ROOM", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conexion, resp)
               return nil
          }

          if(! sala.enSala(emisor.getNombre())){
               resp := NewMensajeAClienteBuilder(respuesta).
               añadirRespuesta("LEAVE_ROOM", "NOT_JOINED", msj.Roomname)
               s.enviaMensajeUsuario(conexion, resp)
               return nil
          }

          sala.eliminarUsuario(emisor.getNombre())
   
          notif := NewMensajeAClienteBuilder(dejaSala).
          añadir("roomname", msj.Roomname).
          añadir("username", emisor.nombre)

          sala.mensajePublico(notif, emisor)
 

     case mensajeDesconectado:
          emisor, registrado := s.conexiones[conexion]
          if !registrado {
          conexion.Desconectar()
          return nil
          }

          s.desconexionUsuario(emisor, conexion)
     
     
}



func (s *servidor) enviaMensajePublico(builder *mensajeAClienteBuilder, cliente *Conexion) {
     for conn := range s.conexiones{
          if (conn != cliente){
               enviaMensajeUsuario(conn, builder)
          }
     }
}


func enviaMensajeUsuario(conexion Conexion, builder *mensajeAClienteBuilder) {
   conexion.enviarMensaje(builder)
}



func (s *servidor) verificaUsuario(nombreUsuario string) bool{
     _, nombre := s.clientes[nombreUsuario]
     return nombre
}


func (s *servidor) desconexionUsuario(usuario *Usuario, con *Conexion) {

     for _, sala := range s.salas {
          users := sala.getUsuarios()
          _, unido := users[usuario.getNombre()]
          if unido{
               sala.eliminarUsuario(usuario.getNombre())        
               notif := NewMensajeAClienteBuilder(dejaSala).
               añadir("roomname", sala.getNombre()).
               añadir("username", usuario.getNombre())

               sala.mensajePublico(notif, usuario)
          }
          sala.eliminaInvitado(usuario.getNombre())
     }
    

     notifDesconectado := NewMensajeAClienteBuilder(desconectado).
     añadir("username", usuario.nombre)

     s.enviaMensajePublico(notifDesconectado, usuario)

     delete(s.clientes, usuario.getNombre())
     delete(s.conexiones, con)
     con.Desconectar()
}




