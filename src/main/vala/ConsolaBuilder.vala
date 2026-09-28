

public class ConsolaBuilder{


    public static string construyeLogin(string usuario){
        var builder = new mensajeAServidorBuilder("IDENTIFY")
        .añadir("username", usuario);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeStatus(string status){
        var builder = new mensajeAServidorBuilder("STATUS")
        .añadir("status", status);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeUsuarios(){
        var builder = new mensajeAServidorBuilder("USERS");

        string msj = builder.construye();
        return msj;
    }

    public static string construyeTexto(string usuario, string mensaje){
        var builder = new mensajeAServidorBuilder("TEXT")
        .añadir("username", usuario)
        .añadir("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyePublico(string mensaje){
        var builder = new mensajeAServidorBuilder("PUBLIC_TEXT")
        .añadir("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeCreaSala(string sala){
        var builder = new mensajeAServidorBuilder("NEW_ROOM")
        .añadir("roomname", sala);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeInvitacion(string sala, string[] usuarios){
        var builder = new mensajeAServidorBuilder("INVITE")
        .añadir("roomname", sala)
        .añadeListaUsuarios(usuarios);

        string msj = builder.construye();
        return msj;
    }
    
    public static string construyeUnirse(string sala){
        var builder = new mensajeAServidorBuilder("JOIN_ROOM")
        .añadir("roomname", sala)

        string msj = builder.construye();
        return msj;
    }

    public static string construyeUsuariosSala(string sala){
        var builder = new mensajeAServidorBuilder("ROOM_USERS")
        .añadir("roomname", sala)

        string msj = builder.construye();
        return msj;
    }

    public static string construyeTextoSala(string sala, string mensaje){
        var builder = new mensajeAServidorBuilder("ROOM_TEXT")
        .añadir("roomname", sala)
        .añadir("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeDejaSala(string sala){
        var builder = new mensajeAServidorBuilder("LEAVE_ROOM")
        .añadir("roomname", sala)

        string msj = builder.construye();
        return msj;
    }
    
    public static string construyeDesconectar(){
        var builder = new mensajeAServidorBuilder("DISCONNECT")

        string msj = builder.construye();
        return msj;
    }





}