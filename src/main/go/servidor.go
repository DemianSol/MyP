// modificación estructura/diseño   --- > agregar al siguiente commit 
// cambio en inicia servidor      ----- > agregar al sigueinte commit 

package main

import (
       "net"   
       "fmt"
       "os"
       "os/signal"
       "syscall"
)
type status string

const(
     AWAY status = "AWAY"
     ACTIVE status = "ACTIVE"
     BUSY status = "BUSY"
)

type servidor struct{
     puerto string
     enchufe net.Listener
     contadorConexiones int
     conexiones map[*conexion]*usuario
     clientes map[string]*usuario // cambiar para manejar clase de Vala
     salas map[string]*sala
     conexionActiva bool
     buzonTareas chan tarea 
}

func newServidor(puerto string) *servidor{
     server := servidor{puerto: puerto}
     server.contadorConexiones = 0
     server.conexiones = make(map[*conexion]*usuario)
     server.salas = make(map[string]*sala)
     server.clientes = make(map[string]*usuario)
     server.buzonTareas = make(chan tarea, 1000) 
     return &server
}

// ESTO FUE TOMADO DE GOBYEXAMPLE/TCP SERVER.COM  
func (s *servidor) iniciaServidor() error{  
     var err error
     s.conexionActiva = true
     fmt.Printf("El puerto %s está disponible\n", s.puerto) 

     s.enchufe, err = net.Listen("tcp", s.puerto)   
     if err != nil{ 
               return fmt.Errorf("Error al escuchar el puerto %s: %w", s.puerto, err) 
          }
     defer s.enchufe.Close()

     // https://gobyexample.com/signals
     // https://leapcell.medium.com/use-chan-os-signal-to-manage-os-signals-in-go-5b0d4d2818fb

     señal := make(chan os.Signal, 1)
	signal.Notify(señal, os.Interrupt, syscall.SIGTERM)

     go func() {
		<-señal
		fmt.Println("\nCerrando servidor y liberando puerto\n")
		s.enchufe.Close() // Cierra el listener TCP inmediatamente
		os.Exit(0)
	}()


     go s.manejaBuzon()
     
     fmt.Printf("Servidor escuchando\n") // creo que está mal   

     for s.conexionActiva{
          conn, err := s.enchufe.Accept()
          if err != nil{
               break
          }
          s.contadorConexiones++
          conex := newConexion(conn, s.contadorConexiones)
     
          fmt.Printf("Conexion recibida de: %d", conex.getIdentificador()) // agregar información 
          
          go s.recibeMensajes(conex) 
     }
     return nil
}


func (s *servidor) imprimeMensaje(mensaje string){ //probablemente está mal
     fmt.Println(mensaje)
}

func (s *servidor) recibeMensajes(con *conexion){
     // que hago cuando se termine? debo mandar mensaje de desconexion? 
     for {
          bytes, err := con.recibeMensaje()
          if err != nil{
               break
          }
          message, err := procesaJSON(bytes)
          if err != nil{
               continue
          }
          s.buzonTareas <- tarea{conn: con, msj: message}
     }
}

// https://blog.joedayz.pe/channels-en-go-buffered-vs-unbuffered-comunicacion-segura-entre-goroutines
func (s *servidor) manejaBuzon() error{ //errores en hilos
     var err error
     for tarea := range s.buzonTareas{
          s.procesaMensaje(tarea.getConexion(), tarea.getMensaje())
     }
     
     if err != nil{
          return fmt.Errorf("Error al procesar la petición: %w", err)
     }
     return nil
}
// https://go.dev/doc/effective_go#type_switch
func (s *servidor) procesaMensaje(conex *conexion, msg mensaje) error{
     if (! conex.getConexionActiva()){
          return nil
     }
     switch msj := msg.(type) {

     case mensajeIdentificar:

          _, user := s.conexiones[conex]
          if (user){
               conex.desconectar()
               return nil // que debo hacer si un usuario ya identificado intenta volver a identificarse?
          }

          if (s.verificaUsuario(msj.Username)){
               resp := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("IDENTIFY", "USER_ALREADY_EXISTS", msj.Username)
               s.enviaMensajeUsuario(conex, resp)
               return nil
          }

          nuevoUsuario := newUsuario(msj.Username, conex)
          s.clientes[msj.Username] = nuevoUsuario
          s.conexiones[conex] = nuevoUsuario
          conex.setAceptado(true)


          res := newMensajeAClienteBuilder(respuesta).
          añadirRespuesta("IDENTIFY", "SUCCESS", msj.Username)
          s.enviaMensajeUsuario(conex, res)

          todos := newMensajeAClienteBuilder(usuarioNuevo).
          añadir("username", msj.Username)
          s.enviaMensajePublico(todos, conex)
               
          

     case mensajeNewStatus: 
          user, b := s.conexiones[conex]
          if !b{
               conex.desconectar()
               return nil 
          }

          estado := status(msj.Status)
          if (estado != ACTIVE && estado != AWAY && estado != BUSY){
               return nil
          }
          if (user.getStatus() != estado){
               user.setStatus(estado)
               builder := newMensajeAClienteBuilder(nuevoStatus).
               añadir("username", user.getNombre()).
               añadir("status", string(estado))
               s.enviaMensajePublico(builder, conex)
          }


     case mensajeListaUsuarios: 
          _, registrado := s.conexiones[conex]
          if !registrado {
            conex.desconectar()
            return nil
          }
          u := make(map[string]string)
          for _, user := range s.clientes{
               u[user.getNombre()] = string(user.getStatus())
          }
          builder := newMensajeAClienteBuilder(listaUsuario).
          añadeListaUsuarios(u)
          s.enviaMensajeUsuario(conex, builder)

     case mensajeTexto:
          emisor, b := s.conexiones[conex]
          if !b{
               conex.desconectar()
               return nil 
          }

          destinatario, user := s.clientes[msj.Username]
          if !user {
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("TEXT", "NO_SUCH_USER", msj.Username)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }


          destino := newMensajeAClienteBuilder(textoDesde).
          añadir("username", emisor.getNombre()).
          añadir("text", msj.Text)
          s.enviaMensajeUsuario(destinatario.getConexion(), destino)
          



     case mensajeTextoPublico:
          remitente, b := s.conexiones[conex]
          if !b {
               conex.desconectar()
               return nil 
          }

          msg := newMensajeAClienteBuilder(textoPublico).
          añadir("username", remitente.getNombre()).
          añadir("text", msj.Text)

          s.enviaMensajePublico(msg, remitente.getConexion())

     case mensajeNuevaSala:
          
          _, existe := s.salas[msj.Roomname]

          if !existe {
               remitente, b := s.conexiones[conex]
               if !b{
                    conex.desconectar()
               }
               room := newSala(msj.Roomname, remitente)
               s.salas[msj.Roomname] = room
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("NEW_ROOM", "SUCCESS", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }

          answer := newMensajeAClienteBuilder(respuesta).
          añadirRespuesta("NEW_ROOM", "ROOM_ALREADY_EXISTS", msj.Roomname)
          s.enviaMensajeUsuario(conex, answer)
          return nil
          
     case mensajeInvita:
          remitente, b := s.conexiones[conex]
          if !b {
               conex.desconectar()
               return nil 
          }


          sala, existe := s.salas[msj.Roomname]

          if(! existe){
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("INVITE", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }
          
          if (! sala.enSala(remitente.getNombre())){
               return nil
           }
          
          for _, nombre := range msj.Users{
               user, answer := s.clientes[nombre]
               if (!answer){
                    respuesta := newMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("INVITE", "NO_SUCH_USER", user.getNombre())
                    s.enviaMensajeUsuario(conex, respuesta)
                    return nil
               }
          }

          invitacion := newMensajeAClienteBuilder(invitacion).
          añadir("username", remitente.getNombre()).
          añadir("roomname", msj.Roomname)

          for _, user := range msj.Users{
               if (sala.enSala(user) || sala.estaInvitado(user)){
                    continue
               }
               sala.invitar(user)
               clientConn, _ := s.clientes[user]
               s.enviaMensajeUsuario(clientConn.getConexion(), invitacion)
          }
          return nil

          
     case mensajeUnirSala:
          remitente, b := s.conexiones[conex]
          if !b{
               conex.desconectar()
               return nil 
          }
          
          sala, existe := s.salas[msj.Roomname]

          if (!existe){
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("JOIN_ROOM", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }

          if (!sala.estaInvitado(remitente.getNombre())){
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("JOIN_ROOM", "NOT_INVITED", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }

          respuesta := newMensajeAClienteBuilder(respuesta).
          añadirRespuesta("JOIN_ROOM", "SUCCESS", msj.Roomname)
          s.enviaMensajeUsuario(conex, respuesta)

          sala.agregarCliente(remitente)

          msjPublico := newMensajeAClienteBuilder(unidoASala).
          añadir("roomname", msj.Roomname).
          añadir("username", remitente.getNombre())

          sala.mensajePublico(msjPublico, remitente)
          return nil

     case mensajeUsuariosSala:
          remitente, b := s.conexiones[conex]
          if !b{
               conex.desconectar()
               return nil 
          }
          sala, existe := s.salas[msj.Roomname]

          if (!existe){
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("ROOM_USRS", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }

          if(!sala.enSala(remitente.getNombre())){
               respuesta := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("ROOM_USERS", "NOT_JOINED", msj.Roomname)
               s.enviaMensajeUsuario(conex, respuesta)
               return nil
          }

          lista := make(map[string]string)
          for _, u := range sala.getUsuarios() {
            lista[u.getNombre()] = string(u.getStatus())
        }
          respuesta := newMensajeAClienteBuilder(usuariosSala).
          añadir("roomname", msj.Roomname).
          añadeListaUsuarios(lista)
          s.enviaMensajeUsuario(conex, respuesta)
          return nil

     case mensajeTextoSala:
          emisor, registrado := s.conexiones[conex]
          if !registrado {
               conex.desconectar()
               return nil 
          }

          sala, existe := s.salas[msj.Roomname]
          if !existe {
               resp := newMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("ROOM_TEXT", "NO_SUCH_ROOM", msj.Roomname)
                    s.enviaMensajeUsuario(conex, resp)
                    return nil
          }

          if (! sala.enSala(emisor.getNombre())){     
               resp := newMensajeAClienteBuilder(respuesta).
                    añadirRespuesta("ROOM_TEXT", "NOT_JOINED", msj.Roomname)
                    s.enviaMensajeUsuario(conex, resp)
                    return nil
          }
          notif := newMensajeAClienteBuilder(textoSala).
               añadir("roomname", msj.Roomname).
               añadir("username", emisor.nombre).
               añadir("text", msj.Text)

          sala.mensajePublico(notif, emisor)

     case mensajeDejaSala:
          emisor, registrado := s.conexiones[conex]
          if !registrado {
               conex.desconectar()
               return nil 
          }

          sala, existe := s.salas[msj.Roomname]
          if !existe {
            resp := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("LEAVE_ROOM", "NO_SUCH_ROOM", msj.Roomname)
               s.enviaMensajeUsuario(conex, resp)
               return nil
          }

          if(! sala.enSala(emisor.getNombre())){
               resp := newMensajeAClienteBuilder(respuesta).
               añadirRespuesta("LEAVE_ROOM", "NOT_JOINED", msj.Roomname)
               s.enviaMensajeUsuario(conex, resp)
               return nil
          }

          sala.eliminarUsuario(emisor.getNombre())
   
          notif := newMensajeAClienteBuilder(dejaSala).
          añadir("roomname", msj.Roomname).
          añadir("username", emisor.nombre)

          sala.mensajePublico(notif, emisor)
 

     case mensajeDesconectado:
          emisor, registrado := s.conexiones[conex]
          if !registrado {
               conex.desconectar()
               return nil
          }

          s.desconexionUsuario(emisor, conex)
     }
     return nil
}



func (s *servidor) enviaMensajePublico(builder *mensajeAClienteBuilder, cliente *conexion) {
     for conn := range s.conexiones{
          if (conn != cliente){
               s.enviaMensajeUsuario(conn, builder)
          }
     }
}


func (s *servidor) enviaMensajeUsuario(conex *conexion, builder *mensajeAClienteBuilder) {
   conex.enviaMensaje(builder)
}



func (s *servidor) verificaUsuario(nombreUsuario string) bool{
     _, nombre := s.clientes[nombreUsuario]
     return nombre
}


func (s *servidor) desconexionUsuario(usuario *usuario, con *conexion) {

     for _, sala := range s.salas {
          users := sala.getUsuarios()
          _, unido := users[usuario.getNombre()]
          if unido{
               sala.eliminarUsuario(usuario.getNombre())        
               notif := newMensajeAClienteBuilder(dejaSala).
               añadir("roomname", sala.getNombre()).
               añadir("username", usuario.getNombre())

               sala.mensajePublico(notif, usuario)
          }
          sala.eliminaInvitado(usuario.getNombre())
     }
    

     notifDesconectado := newMensajeAClienteBuilder(desconectado).
     añadir("username", usuario.nombre)

     s.enviaMensajePublico(notifDesconectado, usuario.getConexion())

     delete(s.clientes, usuario.getNombre())
     delete(s.conexiones, con)
     con.desconectar()
}




