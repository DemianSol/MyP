package main



type usuario struct{
	nombre string
	status status
	conexion *conexion
}


func newUsuario(nombre string, conexion *conexion) *usuario{
	u := new(usuario)
	u.nombre = nombre
	u.conexion = conexion
	u.status = ACTIVE
	return u
}

func (u *usuario) getConexion() *conexion{
	return u.conexion
}

func (u *usuario) getNombre() string{
	return u.nombre
}

func (u *usuario) getStatus() status{
	return u.status
}

func (u *usuario) setStatus(stat status){
	u.status = stat
}