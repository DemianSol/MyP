using GLib;

public class Cliente{
    
    private InetAddress ip;  
    private uint16 puerto;
    private Conexion conexion;
    private bool conexionActiva;

    public Cliente(string ip, string puerto){
    
        this.ip = new InetAddress.from_string(ip);
        this.puerto = (uint16)int.parse(puerto);
        this.conexionActiva = true;
    }


 // https://docs.vala.dev/sample-code/basics/gio-networking-sample.html
    public void iniciaCliente() throws Error{
        

        var cliente = new SocketClient();
        var conn = cliente.connect(new InetSocketAddress (ip, puerto));
        Consola.mostrarMensaje("Conectado al servidor \n");

        this.conexion = new Conexion(conn);

        new Thread<void*>("Hilo servidor", this.recibeMensajesServidor);
        Consola consola = new Consola(this);
        new Thread<void*>("Hilo terminal", consola.leerTerminal);

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

    // https://wiki.gnome.org/Projects/Vala/Tutorial#Dynamic_Type_Casting
    // https://stackoverflow.com/questions/73976240/why-as-keyword-is-not-casting-in-this-code
    private void procesaMensajeServidor(string json_recibido)throws Error {
            try {
                Mensaje? msg = mensajeDeServidorParser.procesarJSON(json_recibido);
                if (msg == null) return;

                if (msg is MensajeNuevoUsuario) {
                    var m = (MensajeNuevoUsuario) msg;
                    Consola.mostrarMensaje("\nNuevo usuario conectado: %s\n", m.usuario);
                }

                else if (msg is MensajeNuevoStatus){
                    var u = (MensajeNuevoStatus) msg;
                    Consola.mostrarMensaje("\nUsuario %s cambio a status: %s", u.usuario, u.stat);
                }

                else if (msg is MensajeListaUsuarios){
                    var m = (MensajeListaUsuarios) msg;
                    Consola.mostrarMensaje(m.listado);
                }

                else if (msg is MensajeTextoDe){
                    var m = (MensajeTextoDe) msg;
                    Consola.mostrarMensaje("%s te envía mensaje: %s", m.usuario, m.texto);
                }

                else if (msg is MensajeTextoPubico){
                    var m = (MensajeTextoPubico) msg;
                    Consola.mostrarMensaje("%s le envía mensaje a todos: %s", m.usuario, m.texto);
                }

                else if (msg is MensajeUnidoASala){
                    var m = (MensajeUnidoASala) msg;
                    Consola.mostrarMensaje("%s se unió a la sala: %s", m.nombre, m.sala);
                }

                else if (msg is MensajeUsuariosSala){
                    var m = (MensajeUsuariosSala) msg;
                    Consola.mostrarMensaje("Usuarios en %s", m.sala);
                    Consola.mostrarMensaje(m.usuarios);
                }

                else if (msg is MensajeTextoSala){
                    var m = (MensajeTextoSala) msg;
                    Consola.mostrarMensaje("(En %s) %s envía mensaje: %s", m.sala, m.nombre, m.texto);
                }

                else if (msg is MensajeDejaSala){
                    var m = (MensajeDejaSala) msg;
                    Consola.mostrarMensaje("(En %s) %s abandonó la sala", m.sala, m.nombre);
                }

                else if (msg is MensajeRespuesta){
                    var m = (MensajeDejaSala) msg;
                    Consola.mostrarMensaje("(En %s) %s abandonó la sala", m.sala, m.nombre);
                }

                else if (msg is MensajeDesconectado){
                    var m = (MensajeDesconectado) msg;
                    Consola.mostrarMensaje("El usuario %s se desconectó", m.nombre);
                }
                else {
                    warning ("El mensaje no es de un tipo aceptado");
                }

            } catch (Error e){
                Consola.mostrarMensaje("Error procesando el JSON: %s\n", e.message);
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


