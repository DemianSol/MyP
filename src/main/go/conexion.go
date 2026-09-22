package main

import (
	"errors"
	"net"
	"bufio"
	"fmt"

)

type conexion struct{

	enchufe net.Conn
	identificador int 
	conexionActiva bool
	lector *bufio.Scanner
	imprimir *bufio.Writer
	aceptado bool
}
func newConexion(enchufe net.Conn, identificador int) *conexion{ 

	return &conexion{
		enchufe: enchufe,
		conexionActiva: true,
		identificador: contadorConexiones+1,
		lector: bufio.NewScanner(enchufe),
		imprimir: bufio.NewWriter(enchufe), 
		aceptado: false,
	}
}

func (c *conexion) setAceptado(v bool){
	c.aceptado = v
}

func (c *conexion) estaAceptador() bool{
	return c.aceptado
}


func (c conexion) getIdentificador() int{
	return c.identificador
}

func (c conexion) getConexionActiva() bool{
	return c.conexionActiva
}

func (c *conexion) desconectar(){
	c.conexionActiva = false
	c.enchufe.Close()

}

func (c *conexion) recibeMensaje() ([]byte, error){
	var bytes []byte
   	
	if c.lector.Scan(){
		bytes, err := c.lector.Bytes()
		if err != nil{
			return fmt.Errorf("Error al leer el mensaje: %w", err)
		}	
	}
	return bytes, nil

}



func (c* conexion) enviaMensaje(builder *mensajeAClienteBuilder) error{

	json, err := builder.construye()
	if err != nil{
		return fmt.Errorf("Error al construir el JSON: %w", err) 
	}

	jsonSalto := append([]byte(json), '\n')
	porEnviar, err := c.imprimir.Write(jsonSalto)
	if err != nil{
		return fmt.Errorf("Error al enviar por el socket: %w", err)
	}

	err := c.imprimir.Flush()
	if err != nil{
		return fmt.Errorf("Error en flush para el socket: %w", err) 
	}
	return nil
}


