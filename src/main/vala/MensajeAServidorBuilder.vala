using Json;
using GLib;

public class MensajeAServidorBuilder : GLib.Object {
	private Json.@Object raiz;

	
	public MensajeAServidorBuilder(string tipoMensaje){
		this.raiz = new Json.Object();
		this.raiz.set_string_member("type", tipoMensaje);
	}


	public MensajeAServidorBuilder agregar(string llave, string valor){
		this.raiz.set_string_member(llave, valor);
		return this;
	}

	public MensajeAServidorBuilder agregaListaUsuarios(string[] usuarios){
		var users = new Json.Array();
		foreach (string u in usuarios){
			users.add_string_element(u);
		}
		this.raiz.set_array_member("users", users);
		return this;
	}

	public string construye(){
		var nodo = new Json.Node(NodeType.OBJECT);
		nodo.set_object(this.raiz);
		var genera = new Json.Generator();
		genera.set_root(nodo);
		return genera.to_data(null) + "\n";
	}

}


