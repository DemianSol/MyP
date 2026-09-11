public class Cliente : Object{

    private string nombreUsuario;
    private HashMap<Sala> salasActivas;
    private Status status;
    private Conexion conexion;    // ´rebed otséste método debería importarse del código en el otro lenguaje
}


public Cliente(string nombreUsuario, Conexion conexion){
    this.nombreUsuario = nombreUsuario;
    this.conexion = conexion;
}


public boolean solicitarConexion(){


}


public boolean cerrarConexion(){

}


private void recibeMensaje(){


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




