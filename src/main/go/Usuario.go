

type Status string

const (
	ACTIVE Status = "ACTIVE"
	BUSY Status = "BUSY"
	AWAY Status = "AWAY"
)

type Usuario struct{
	nombre string
	status Status
	conexion *Conexion
}


func NewUsuario(nombre string, conexion *Conexion) *Usuario{
	u := new(Usuario)
	u.nombre = nombre
	u.conexion = conexion
	u.status = ACTIVE
	return u
}

func (u *Usuario) getConexion() *Conexion{
	return u.conexion
}

func (u *Usuario) getNombre() string{
	return u.nombre
}

func (u *Usuario) getStatus() Status{
	return u.status
}

func (u *Usuario) setStatus(status Status) error{
	u.status = status
}