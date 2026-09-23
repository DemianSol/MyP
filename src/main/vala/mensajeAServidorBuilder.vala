using Json;

public class mensajeAServidorBuilder : Object {
	private JSON.Object raiz;

	// https://aztlan.fciencias.unam.mx/gitlab/canek/gtrophies/-/blob/9ec8ba9743fb58d2a3a77ee2664d86b6632517bf/lib/psn/translator.vala
	// https://valadoc.org/json-glib-1.0/Json.Object.set_string_member.html
	public mensajeAServidorBuilder(string tipoMensaje){
		this.raiz = new Json.Object();
		this.raiz.set_string_member = ("type", tipoMensaje);
	}


	public unowned mensajeAServidorBuilder añadir(string llave, string valor){
		this.raiz.set_string_member(llave, valor);
		return this;
	}

	// https://docs.vala.dev/sample-code/other/json-sample.html
	public unowned mensajeAServidorBuilder añadeListaUsuarios(string[] usuarios){
		var users = new Json.Array()
		foreach (string u in usuarios){
			users.add_string_element(u);
		}
		this.raiz.set_array_member("users", users);
		return this;
	}

	public string contruye(){
		var nodo = new Json.Node(NodeType.Object);
		nodo.set_object(this.raiz);
		var genera = new Json.Generator();
		genera.set_root(nodo);
		return genera.to_data(null) + "\n";
	}

}


