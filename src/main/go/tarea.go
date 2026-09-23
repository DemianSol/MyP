package main


type tarea struct{
	conn *conexion
	msj mensaje
}

func newTarea(conn *conexion, message mensaje) *tarea{
	return &tarea{
		conn : conn,
		msj : message,
	}
}

func (t tarea) getConexion() *conexion{
	return t.conn
}


func (t tarea) getMensaje() mensaje{
	return t.msj
}
