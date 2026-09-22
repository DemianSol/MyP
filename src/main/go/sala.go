

type sala struct{
	nombreSala string
	usuariosSala map[string]*usuario
	invitados map[string]bool
}

func newSala(nombreSala string, creador *usuario) *sala{
	s := &sala{ 
		nombreSala: nombreSala,
		usuariosSala: make(map[string]*usuario)
		invitados: make(map[string]bool),
	}

	s.usuariosSala[creador.getNombre()] = creador
	s.invitados[creador.getNombre()] = true

	return s
}

func (s *sala) getNombre() string{
	return s.nombreSala
}


func (s *sala) getUsuarios() (map[string]*usuario){
	return s.usuariosSala
}

func (s *sala) enSala(user string) bool {
	u, b := s.usuariosSala[user]
	return b
}

func (s *sala) estaInvitado(user string) bool{
	return s.invitados[user]
}

func (s *sala) invitar(clienteNuevo string){
	s.invitados[clienteNuevo] = true
}


func (s *sala) agregarCliente(user *usuario){
	s.usuariosSala[user.getNombre()] = user
	delete(s.invitados, user.getNombre())
}

func (s *sala) eliminarUsuario(user string){
	delete(s.usuariosSala, user)
}

func (s *sala) eliminaInvitado(user string){
	delete(s.invitados, user)
}

func (s *sala) mensajeGeneral(builder *mensajeAClienteBuilder){
	for _, u := range s.usuariosSala{
		u.getConexion().enviarMensaje(builder)
	}
}

func (s *sala) mensajePublico(builder *mensajeAClienteBuilder, user *usuario){
	for _, u := range s.usuariosSala{
		if (u != user){
			u.getConexion().enviarMensaje(builder)
		}
	}
}




