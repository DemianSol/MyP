

type tarea struct{
	conn conexion
	mensaje Mensaje


func newTarea(conn conexion, mensaje Mensaje){
	return &tarea{
		conn : conn
		mensaje : mensaje
	}
}

func (t tarea) getConexion(){
	return t.conexion
}


func (t tarea) getMensaje(){
	return t.mensaje
}
}