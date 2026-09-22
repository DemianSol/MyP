

type tarea struct{
	conn *conexion
	msj mensaje
}

func newTarea(conn *conexion, msj mensaje){
	return &tarea{
		conn : conn
		mensaje : mensaje
	}
}

func (t tarea) getConexion(){
	return t.conn
}


func (t tarea) getMensaje(){
	return t.msj
}
