using GLib

class Cliente{

    private string nombreUsuario;
    private HashMap<Sala> salasActivas;
    private Status status;
    private Conexion conexion;    // ´rebed otséste método debería importarse del código en el otro lenguaje
}


public Cliente(string nombreUsuario, Conexion socket){
    this.nombreUsuario = nombreUsuario;
}


public void conecta(){

    int puerto = 0 // recibe puerto del cliente, entonces usa a Conexion 
    try{
        this.socket = new SocketClient();
        this.conexion = new Conexion(socket);    
        var hilo = new Thread<*void> (null, conexion.recibeMensaje())
    }
    
}


public boolean cerrarConexion(){

}


private void recibeMensaje(Mensaje mensaje){


}



public void invitarASala(Cliente cliente){

}



public void enviarMensajeTodos(string mensaje){


}


public void enviarMensajeSala(string mensaje){

    
}



public void enviarMensajeUsuario(string mensaje){

    
}


public void cambioStatus(Status status){


}




