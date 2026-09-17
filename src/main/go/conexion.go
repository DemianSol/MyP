import (
	"erros"
	"net"
	"io"
	"bufio"
	"os"
	"fmt"

)

type Conexion struct{

	enchufe net.Conn
	buzon chan tarea
	identificador int 
	conexionActiva bool
	contadorConexiones int
	lector bufio
	imprimir bufio.Writer

func NewConexion(enchufe net.Dial, buzon chan<- tarea) *conexion{ //solo puede enviar a la cola

	return &conexion{
		enchufe: enchufe
		buzon: buzon
		conexionActiva: true
		contadorConexiones: identificador+1
		identificador: contadorConexiones+1
		lector: bufio.NewScanner(os.Stdin)
		imprimir: bufio.Writer(enchufe) 
	}
}


func (c conexion) getIdentificador(){
	return c.dentificador
}

func (c conexion) getConexionActiva(){
	return c.conexionActiva
}

func (c *conexion) desconectar(){

	c.conexionActiva = false
	return c.enchufe.Close()

}

func (c conexion) recibeMensaje() error{
	var err error
     
	while (conexionActiva){
		if c.lector.Scan(){
			texto := lectura.Text()
			MensajeCliente m, error:= MensajeCliente.getMensaje(texto)
			if err != nil{
				fmt.Errof("Error al procesar mensaje: %w", err)
			}
			buzon <- tarea{conexion, m}
		}
		if err != nil{
			fmt.Error("El mensaje no se leyó correctamente: %w", err)
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