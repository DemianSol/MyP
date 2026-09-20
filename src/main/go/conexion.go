import (
	"error"
	"net"
	"io"
	"bufio"
	"os"
	"fmt"

)

type Conexion struct{

	enchufe net.Conn
	identificador int 
	conexionActiva bool
	contadorConexiones int
	lector bufio
	imprimir bufio.Writer

func NewConexion(enchufe net.Conn) *conexion{ //solo puede enviar a la cola

	return &conexion{
		enchufe: enchufe
		conexionActiva: true
		contadorConexiones: identificador+1
		identificador: contadorConexiones+1
		lector: bufio.NewScanner(os.Stdin)
		imprimir: bufio.Writer(enchufe) 
	}
}


func (c conexion) getIdentificador() int{
	return c.dentificador
}

func (c conexion) getConexionActiva() bool{
	return c.conexionActiva
}

func (c *conexion) desconectar(){
	c.conexionActiva = false
	return c.enchufe.Close()

}

func (c conexion) recibeMensaje() ([]byte, error){
	var err error
     
	while (conexionActiva){
		if c.lector.Scan(){
			bytes := lectura.Bytes()
			if err != nil{
				fmt.Errorf("Error al leer el mensaje: %w", err)
			}
			return bytes, nil
		}
	}
}



func (c* conexion) enviaMensaje(mensaje Mensaje, err error){

	m, err := c.imprimir.WriteString(mensaje.toString())
	
	if err != nil{
		// algo para enviar error a los susarios 
	}
}


}