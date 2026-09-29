
public class Consola{

    Cliente cliente;
    private bool activo;
    private static Mutex mutex = Mutex();

    public Consola(Cliente cliente){
        this.cliente = cliente;
        this.activo = true;

    }


    public void* leerTerminal () {
        while (this.activo) {
            mutex.lock ();
            print ("> ");
            stdout.flush ();
            mutex.unlock ();

            string? entrada = stdin.read_line ();

            if (entrada == null) {
                this.activo = false;
                this.cliente.desconectar();
                break;
            }

            string texto = entrada.strip();

            if (texto == null){
                continue;
            }

            if (texto == "/desconectar"){
                this.activo = false;
                this.cliente.desconectar();
                break;
            }

            var msj = ConsolaParser.parsearConsola(texto);
            if (msj != null){
                this.cliente.enviar(msj);
            }
        }
        return null;
    }

    public static void mostrarMensaje(string texto) {
        mutex.lock();
        try{
            print("\r\033[K%s\n> ", texto);
        } finally {
            mutex.unlock();
        }
    }

    public static void mostrarLista(string[] mensajes) {
        mutex.lock();
        try{
            string resultado = string.joinv (", ", mensajes);
            print("[%s]\n", resultado);
        } finally{
            mutex.unlock();
        }
    }


    public void desconectar(){
        this.activo = false;    
    }

}