

public class ConsolaBuilder{


    public static string construyeLogin(string usuario){
        var builder = new MensajeAServidorBuilder("IDENTIFY")
        .agregar("username", usuario);

        string msj = builder.construye();
        Consola.mostrarMensaje(msj);
        return msj;
    }

    public static string construyeStatus(string status){
        var builder = new MensajeAServidorBuilder("STATUS")
        .agregar("status", status);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeUsuarios(){
        var builder = new MensajeAServidorBuilder("USERS");

        string msj = builder.construye();
        return msj;
    }

    public static string construyeTexto(string usuario, string mensaje){
        var builder = new MensajeAServidorBuilder("TEXT")
        .agregar("username", usuario)
        .agregar("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyePublico(string mensaje){
        var builder = new MensajeAServidorBuilder("PUBLIC_TEXT")
        .agregar("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeCreaSala(string sala){
        var builder = new MensajeAServidorBuilder("NEW_ROOM")
        .agregar("roomname", sala);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeInvitacion(string sala, string[] usuarios){
        var builder = new MensajeAServidorBuilder("INVITE")
        .agregar("roomname", sala)
        .agregaListaUsuarios(usuarios);

        string msj = builder.construye();
        return msj;
    }
    
    public static string construyeUnirse(string sala){
        var builder = new MensajeAServidorBuilder("JOIN_ROOM")
        .agregar("roomname", sala);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeUsuariosSala(string sala){
        var builder = new MensajeAServidorBuilder("ROOM_USERS")
        .agregar("roomname", sala);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeTextoSala(string sala, string mensaje){
        var builder = new MensajeAServidorBuilder("ROOM_TEXT")
        .agregar("roomname", sala)
        .agregar("text", mensaje);

        string msj = builder.construye();
        return msj;
    }

    public static string construyeDejaSala(string sala){
        var builder = new MensajeAServidorBuilder("LEAVE_ROOM")
        .agregar("roomname", sala);

        string msj = builder.construye();
        return msj;
    }
    
    public static string construyeDesconectar(){
        var builder = new MensajeAServidorBuilder("DISCONNECT");

        string msj = builder.construye();
        return msj;
    }





}