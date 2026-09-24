
// https://docs.vala.dev/tutorials/main/03-00-object-oriented-programming/03-12-dynamic-type-casting.html
// https://docs.vala.dev/tutorials/main/03-00-object-oriented-programming/03-08-interfaces.html
public interface Mensaje : Object {
    // public abstract string tipoMensaje(); no estoy seguro de si es necesario declarar metodo abstracto
}


public class MensajeNuevoUsuario : Object, Mensaje {
    public string usuario{ get; set;}

    public MensajeNuevoUsuario(string usuario){
        this.usuario = usuario;
    }

  /* public string tipoMensaje(){
        public string get_message_type () {
        return "NEW_USER";     EJEMPLO DE COMO SE VERÍA SI DECLARO METODO ABSTRACTO
    }*/
    
}


public class MensajeNuevoStatus : Object, Mensaje {
    public string usuario{ get; set;}
    public string stat{get; set;}

    public MensajeNuevoStatus(string usuario, string stat){
        this.usuario = usuario;
        this.stat = stat;
    }
}

public class MensajeListaUsuarios : Object, Mensaje{
    public string[] listado;
    
    public MensajeListaUsuarios(string[] listado){
        this.listado = listado;
    } 
}

public class MensajeTextoDe : Object, Mensaje{
    public string usuario{get; set;}
    public string texto{get; set;}

    public MensajeTextoDe(string usuario, string texto){
        this.usuario = usuario;
        this.texto = texto;
    }
}


public class MensajeTextoPubico : Object, Mensaje{
    public string usuario{get; set};
    public string texto{get; set;}

    public MensajeTextoPubico(string usuario, string texto){
        this.usuario = usuario;
        this.texto = texto;
    }
}

public class MensajeUnidoASala : Object, Mensaje{
    public string sala{get; set;};
    public string nombre {get;set;}

    public MensajeUnidoASala(string sala, string nombre){
        this.sala = sala;
        this.nombre = nombre;
    }
}


public class MensajeUsuariosSala : Object, Mensaje{
    public string sala{get; set;};
    public string[] usuarios{get; set;}

    public MensajeUsuariosSala(string sala, string[] usuarios){
        this.sala = sala;
        this.usuarios = usuarios;
    }
}

public class MensajeTextoSala : Object, Mensaje{
    public string sala{get; set;};
    public string nombre{get; set;};
    public string texto{get; set;};

    public MensajeTextoSala(string sala, string nombre, string texto){
        this.sala = sala;
        this.nombre = nombre;
        this.texto = texto;
    }
}

public class MensajeDejaSala : Object, Mensaje{
    public string sala{get; set;};
    public string nombre{get; set;}:

    public MensajeDejaSala(string sala, string nombre){
        this.sala = sala;
        this.nombre = nombre;
    }
}

public class MensajeRespuesta : Object, Mensaje{
    public string operacion{get; set;}:
    public string resultado{get; set;}:
    public string extra{get; set;};

   
    public MensajeRespuesta(string operacion, string resultado, string extra){
        this.operacion = operacion;
        this.resultado = resultado;
        this.extra = extra;
        
    }
}


public class MensajeDesconectado : Object, Mensaje{
    public string nombre{get; set;};

    public MensajeDesconectado(string nombre){
        this.nombre = nombre;
    }
}






