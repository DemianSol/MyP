

type Sala struct{
	nombreSala string
	usuariosSala map[string]*Usuario
	invitados map[string]bool
}

func NewSala(nombreSala string, creador *Usuario) *Sala{
	s := &Sala{ 
		nombreSala: nombreSala,
		usuariosSala: make(map[string]*Usuario)
		invitados: make(map[string]bool),
	}

	s.usuariosSala[creador.getNombre()] = creador
	s.invitados[creador.getNombre()] = true

	return s
}

func (s *Sala) getNombre() string{
	return s.nombreSala
}


func (s *Sala) getUsuarios() (map[string]*Usuario){
	return s.usuariosSala
}

func (s *Sala) enSala(user string) bool {
	u, b := s.usuariosSala[user]
	return b
}

func (s *Sala) estaInvitado(user string) bool{
	return s.invitados[user]
}

func (s *Sala) invitar(clienteNuevo string){
	s.invitados[clienteNuevo] = true
}


func (s *Sala) agregarCliente(user *Usuario){
	s.usuariosSala[user.getNombre()] = user
	delete(s.invitados, user.getNombre())
}

func (s *Sala) eliminarUsuario(user string){
	delete(s.usuariosSala, user)
}

func (s *Sala) eliminaInvitado(user string){
	delete(s.invitados, user)
}

func (s *Sala) mensajeGeneral(builder *mensajeAClienteBuilder){
	for _, u := range s.usuariosSala{
		u.getConexion().enviarMensaje(builder)
	}
}

func (s *Sala) mensajePublico(builder *mensajeAClienteBuilder, user *Usuario){
	for _, u := range s.usuariosSala{
		if (u != user){
			u.getConexion().enviarMensaje(builder)
		}
	}
}




