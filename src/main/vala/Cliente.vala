using GLib;

public class Cliente{
    
    private InetAddress ip;  
    private uint16 puerto;
    private Conexion conexion;
    private bool conexionActiva;
    private MainLoop loop;

    public Cliente(string ip, string puerto){
 
        this.puerto = (uint16)int.parse(puerto);
        this.ip = new InetAddress.from_string(ip);
        this.conexionActiva = false;
        this.loop = new MainLoop();
    }


    public void iniciaCliente() throws Error{
        

        var cliente = new SocketClient();
        var conn = cliente.connect(new InetSocketAddress (ip, puerto));
        Consola.mostrarMensaje("Conectado al servidor \n");

        this.conexion = new Conexion(conn);
        this.conexionActiva = true;

        new Thread<void*>("Hilo servidor", this.recibeMensajesServidor);
        Consola consola = new Consola(this);
        new Thread<void*>("Hilo terminal", consola.leerTerminal);
        this.loop.run(); 

    }


    private void* recibeMensajesServidor(){

        while(this.conexionActiva){
            try{
                string? msj = this.conexion.recibeMensaje();

                if(msj == null){
                    Consola.mostrarMensaje("Cliente desconectado\n");
                    this.conexionActiva = false;
                    break;
                }                
                procesaMensajeServidor(msj);
            } catch(Error e){
                if (this.conexionActiva){
                stderr.printf ("Error al recibir el mensaje: %s\n", e.message);
            }
            break;
            }
        }
        return null;
    
    }
    // 

    private void procesaMensajeServidor(string json_recibido)throws Error {
            try {
                Mensaje? msg = MensajeDeServidorParser.procesarJSON(json_recibido);
                if (msg == null){
                    return;
                }

                if (msg is MensajeNuevoUsuario) {
                    var m = (MensajeNuevoUsuario) msg;
                    Consola.mostrarMensaje(@"\nNuevo usuario conectado: $(m.usuario)\n");
                }

                else if (msg is MensajeNuevoStatus){
                    var u = (MensajeNuevoStatus) msg;
                    Consola.mostrarMensaje(@"\nUsuario $(u.usuario) cambio a status: $(u.stat)\n");
                }

                else if (msg is MensajeListaUsuarios){
                    var m = (MensajeListaUsuarios) msg;
                    Consola.mostrarLista(m.listado);
                }

                else if (msg is MensajeTextoDe){
                    var m = (MensajeTextoDe) msg;
                    Consola.mostrarMensaje(@"\n$(m.usuario) te envía mensaje: $(m.texto)\n");
                }

                else if (msg is MensajeTextoPubico){
                    var m = (MensajeTextoPubico) msg;
                    Consola.mostrarMensaje(@"\n$(m.usuario) le envía mensaje a todos: $(m.texto)\n");
                }

                else if (msg is MensajeUnidoASala){
                    var m = (MensajeUnidoASala)msg;
                    Consola.mostrarMensaje(@"\n$(m.nombre) se unió a la sala: $(m.sala)\n");
                }

                else if (msg is MensajeUsuariosSala){
                    var m = (MensajeUsuariosSala)msg;
                    Consola.mostrarMensaje(@"\nUsuarios en $(m.sala)\n");
                    Consola.mostrarLista(m.usuarios);
                }

                else if (msg is MensajeTextoSala){
                    var m = (MensajeTextoSala)msg;
                    Consola.mostrarMensaje(@"\n(En $(m.sala)) $(m.nombre) envía mensaje: $(m.texto)\n");
                }

                else if (msg is MensajeDejaSala){
                    var m = (MensajeDejaSala)msg;
                    Consola.mostrarMensaje(@"\n(En $(m.sala)) $(m.nombre) abandonó la sala\n");
                }

                else if (msg is MensajeRespuesta){
                    var m = (MensajeRespuesta)msg;
                    Consola.mostrarMensaje(@"\nRespuesta recibida por operación $(m.operacion) con resultado $(m.resultado). $(m.extra)\n");
                }

                else if (msg is MensajeDesconectado){
                    var m = (MensajeDesconectado) msg;
                    Consola.mostrarMensaje(@"El usuario $(m.nombre) se desconectó\n");
                }

                else if (msg is MensajeInvitacion){
                    var s = (MensajeInvitacion)msg;
                    Consola.mostrarMensaje(@"\nEl usuario $(s.usuario) te invitó a la sala $(s.sala)\n");
                }

                else {
                    warning ("El mensaje no es de un tipo aceptado");
                }

            } catch (Error e){
                Consola.mostrarMensaje(@"Error procesando el JSON: $(e.message)\n");
            }
    }


    public void enviar(string texto) {
        try {
            this.conexion.enviaMensajeServidor(texto);
        } catch (Error e) {
            Consola.mostrarMensaje("Error al enviar mensaje: " + e.message);
        }
    }

    public void desconectar () {
        this.conexionActiva = false;
        if (this.conexion != null) {
            this.conexion.cerrar ();
        }
    }

}


