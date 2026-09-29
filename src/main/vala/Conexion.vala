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
    
    

    public void enviaMensajeServidor(string msj) throws Error{
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