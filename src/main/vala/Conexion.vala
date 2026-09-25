using GLib;

public class Conexion{

    SocketConnection enchufe;
    DataInputStream input;
    DataOutputStream output;
    
    public Conexion(SocketConnection enchufe){
        this.enchufe = enchufe;
        this.input = new DataInputStream(enchufe.get_input_stream());
        this.input.set_newline_type(DataStreamNewlineType.LF);
        this.output = new DataOutputStream(enchufe.get_output_stream());
    }

    public string? recibeMensaje() throws Error{

        size_t longitud;
        return this.input.read_line(out longitud, null);
    }
    
    
    // https://wiki.gnome.org/Projects(2f)Vala(2f)GIONetworkingSample.html
    // https://docs.vala.dev/genie/sample-code/gio-networking-sample
    private void enviaMensajeServidor(string msj) throws Error{
        try{
            this.output.put_string(msj + "\n");
        } catch(Error e){
            stderr.printf("Error al enviar el mensaje al servidor: %s\n", e.message);
        }
    }

    public void cerrar() {
        try {
            this.enchufe.close();
        } catch (Error e) {}
    }
}