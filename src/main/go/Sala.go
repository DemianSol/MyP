

type sala struct{
	nombreSala string
	usuariosSala map[string]cliente      // el int puede ser la dirección IP del usuario
}

func NewSala(nombreSala string) *sala{
	s := sala{nombreSala: nombreSala}
	s.usuariosSala = make(map[string]cliente)
	return &s
}


func invitarCliente(clienteNuevo cliente){

}


func agregarCliente(clienteNuevo cliente){

}

func usuariosConectados() map[string]cliente{

} 


func abandonaSala(cl cliente){

}


