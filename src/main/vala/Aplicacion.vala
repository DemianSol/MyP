using GLib;

public class Aplicacion{

    private static void errorPuerto(string puerto){
        stderr.printf("Error: El puerto %s es inválido\n", puerto);
        GLib.Process.exit(1);
    }   

    private static void errorDireccion(string direccion){
        stderr.printf("Error: La dirección %s es inválida\n", direccion);
        GLib.Process.exit(1);
    }

    private static void uso(){
        stderr.printf("Uso: ./proyecto1 <direcciónIP> <puerto>\n");
        GLib.Process.exit(1);
    }



    string[] args;      

    public Aplicacion(string[] args){
        if (args.length != 2){
           uso();
       }
       if (! verificaDireccion(args[0])){
            errorDireccion(args[0]);
       }
       if (! verificaPuerto(args[1])){
            errorPuerto(args[1]);
       }
        this.args = args;
    }

    
    private bool verificaDireccion(string ip){
        var direccion = new InetAddress.from_string(ip);
        return direccion != null;
    }


    private bool verificaPuerto(string puerto){
        int64 resultado;

         if (! int64.try_parse(puerto, out resultado)) {
            return false;
         }

         if (resultado < 1025 || resultado > 65535){
            return false;
         }

         return true;

    }

    public void ejecuta(){
        var cliente = new Cliente(this.args[0], this.args[1].to_string());
        try{
            cliente.iniciaCliente();
        } catch(Error e){
            stderr.printf ("No se pudo iniciar el cliente: %s\n", e.message);
        }
    }

}