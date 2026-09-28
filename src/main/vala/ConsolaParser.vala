
// https://valadoc.org/glib-2.0/string.down.html
public class ConsolaParser{

    public static string? parsearConsola(string linea){
        string cadena = linea.strip();
        string[] campos = cadena.split(" ");

        string tipo = campos[0];
        if (! tipo.has_prefix("/")){
            Consola.mostrarMensaje("Uso: </opcion> <campo1> campo<2> ...");
            return null;
        }

        switch(tipo) {

            case "/conectarse":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /conectarse <nombre_usuario>");
                    return null;
                }
                return ConsolaBuilder.construyeLogin(campos[1]);

            case "/estatus":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /estatus <AWAY|BUSY|ACTIVE");
                    return null;
                }
                return ConsolaBuilder.construyeStatus(campos[1]);
            
            case "/usuarios":
                if (campos.length != 1){
                    Consola.mostrarMensaje("Uso: /usuarios");
                    return null;
                }
                return ConsolaBuilder.construyeUsuarios();

            case "/privado":
                if (campos.length < 3){
                    Consola.mostrarMensaje("Uso: /privado <usuario> <mensaje>");
                    return null;
                }
                string usuario = campos[1];
                int tamano = campos[0].length + campos[1].length + 2;
                string msj = cadena.substring(tamano);
                return ConsolaBuilder.construyeTexto(usuario, msj);

            case "/publico":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /publico <mensaje>");
                    return null;
                }
                int tamano = campos[0].length + 1;
                string msj = cadena.substring(tamano);
                return ConsolaBuilder.construyePublico(msj);
            
            case "/creasala":
                if (campos.length !=2){
                    Consola.mostrarMensaje("Uso: /creasala <nombre_sala>");
                    return null;
                }

                return ConsolaBuilder.construyeCreaSala(campos[1]);
            
            case "/invitar":
                if (campos.length < 3){
                    Consola.mostrarMensaje("Uso: /invitar <nombre_sala> <usuario1,usario2,usuario3,...>");
                    return null;
                }
                string sala = campos[1];
                string[] usuarios = campos[2].split(",");

                return ConsolaBuilder.construyeInvitacion(sala, usuarios);
            
            case "/unirse":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /unirse <nombre_sala>");
                    return null;
                }
                string sala = campos[1];
                return ConsolaBuilder.construyeUnirse(sala);

            case "/usuariosala":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /usuariosala <nombre_sala>");
                    return null;
                }

                return ConsolaBuilder.construyeUsuariosSala(campos[1]);

            case "/sala":
                if (campos.length < 3){
                    Consola.mostrarMensaje("Uso: /sala <sala> <mensaje>");
                    return null;
                }
                string sala = campos[1];
                int tamano = campos[0].length + campos[1].length + 2;
                string msj = cadena.substring(tamano);
                return ConsolaBuilder.construyeTextoSala(sala, msj);

            case "/dejasala":
                if (campos.length != 2){
                    Consola.mostrarMensaje("Uso: /dejasala <sala>");
                    return null;
                }
                return ConsolaBuilder.construyeDejaSala(campos[1]);

            case "/desconectar":
                if (campos.length != 1){
                    Consola.mostrarMensaje("Uso: /desconectar");
                    return null;
                }
                return ConsolaBuilder.construyeDesconectar();

            case "/ayuda":
                mostrarAyuda();
                return null;
            
            default:
                Consola.mostrarMensaje("Comando desconocido. Escribe /ayuda para ver cuáles están disponibles.");
                return null;
        }
    }



    private static void mostrarAyuda () {
        string[] ayuda = {
            "--- COMANDOS DISPONIBLES ---",
            "/conectarse <usuario>        - Identificarse en el servidor",
            "/estatus <estado>            - Cambiar tu estado a AWAY, BUSY o ACTIVE",
            "/usuarios                    - Ver lista global de usuarios",
            "/privado <usuario> <mensaje> - Enviar mensaje privado",
            "/publico <mensaje>           - Mensaje a todos",
            "/creasala <nombre>           - Crear nueva sala",
            "/unirse <nombre>             - Entrar a una sala existente",
            "/invitar <sala> <u1,u2,..>   - Invitar usuarios a tu sala",
            "/usuariosala <sala>          - Ver miembros de una sala",
            "/sala <sala> <mensaje>       - Enviar mensaje a una sala",
            "/dejasala <sala>             - Salir de una sala",
            "/desconectar                 - Salir del chat",
        };
        Consola.mostrarLista(ayuda);
    }
}