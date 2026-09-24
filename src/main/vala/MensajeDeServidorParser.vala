using Json;


public class mensajeDeServidorParser : Object {

	// https://docs.vala.dev/sample-code/other/json-sample.html
	public static Mensaje? procesarJSON (string json) throw Error{
		var parser = new Json.Parser();
		parser.load_from_data(json, -1);

		var base = 	parser.get_root().get_object();
		if (base == null || base.get_node_type() != Json.NodeType.OBJECT){
			throw new IOError.INVALID_DATA("El formato JSON no es correcto");
		}

		var msj = base.get_object();
		if (! msj.has_member("type")){
			throw new IOError.INVALID_DATA("El mensaje no contiene el campo 'type'");
		}

		string tipo = msj.get_string_member("type");

		switch(tipo){

			case "NEW_USER":
				string user = msj.get_string_member("username");
            	return new MensajeNuevoUsuario(user);

			case "NEW_STATUS":
				string user = msj.get_string_member("username");
				string stat = msj.get_string_member("status");
				return new MensajeNuevoStatus(user, stat);

			case "USER_LIST":	
				Json.Array listado = msj.get_array_members("users");
				string[] usuarios = {};
				foreach(u in listado.get_elements()){
					usuarios += u.get_string()
				}
				return new MensajeListaUsuarios(usuarios);

			case "TEXTO_FROM":
				string user = msj.get_string_member("usarname");
				string texto = msj.get_string_member("text");
				return new MensajeTextoDe(user, texto);

			case "PUBLIC_TEXT_FROM":
				string user = msj.get_string_member("usarname");
				string texto = msj.get_string_member("text");
				return new MensajeTextoPubico(user, texto);

			case "JOINED_ROOM":
				string sala = msj.get_string_member("roomname");
				string nombre = msj.get_string_member("username");
				return new MensajeUnidoASala(sala, nombre);

			case "ROOM_USER_LIST":
				string sala = msj.get_string_member("roomname");
				Json.Array listado = msj.get_array_members("users");
				string[] usuarios = {};
				foreach(u in listado.get_elements()){
					usuarios += u.get_string()
				}
				return new MensajeUsuariosSala(sala, listado);

			case "ROOM_TEXT_FROM":
				string sala = msj.get_string_member("roomname");
				string nombre = msj.get_string_member("username");
				string texto = msj.get_string_member("text")
				return new MensajeTextoSala(sala, nombre, texto);

			case "LEFT_ROOM":
				string sala = msj.get_string_member("roomname");
				string nombre = msj.get_string_member("username");
				return new MensajeDejaSala(sala, nombre);

			case "RESPONSE":
				string operacion = msj.get_string_member("operation");
				string resultado = msj.get_string_member("result")
				string extra? = null;
				if (msj.has_member("extra")){
					msj.get_string_member("extra")
				}
				return new MensajeRespuesta(operacion, resultado, extra);

			case "DISCONNECTED":
				string usuario = msj.get_string_member("username");
				return new MensajeDesconectado(usuario);
				
			
			default:
				throw new IOError.INVALID_DATA ("Tipo de mensaje es desconocido: %s", tipo);
				return null
		}
	}
}