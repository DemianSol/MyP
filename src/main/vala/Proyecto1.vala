

// https://www.baeldung.com/java-read-input-until-condition

public class Proyecto1 {

    private bool nombreRecibo = false

    public static void main(string[] args){

        Aplicacion aplicacion = new Aplicacion()
        stdout.printf("Como primer mensaje debes escribir tu nombre. Consulta uso() para las opciones. (Para salir usa Ctrl + D): \n"); // chechar

        string? linea = null
        while ((linea = stdin.read_line()) != null) {
            if (!nombre){
                aplicacion.recibeNombre(linea);
                nombreRecibido = true;
            }
            aplicacion.ejectuaMensaje(linea);
            }    
    }
}   
