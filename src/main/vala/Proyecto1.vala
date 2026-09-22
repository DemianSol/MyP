

// https://www.baeldung.com/java-read-input-until-condition

public class Proyecto1 {


    public static void main(string[] args){

        Aplicacion aplicacion = new Aplicacion()
        stdout.printf("Escribe texto (Para salir usa Ctrl + D): \n"); // chechar

        string? linea = null
        while ((linea = stdin.read_line()) != null) {
            aplicacion.ejectuaMensaje(linea);
        }    
    }
}   
