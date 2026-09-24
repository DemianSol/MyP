package main

import (
	"net"
	"bufio"
	"fmt"
	"sync"
)

type conexion struct{

	enchufe net.Conn
	identificador int 
	conexionActiva bool
	lector *bufio.Scanner
	imprimir *bufio.Writer
	aceptado bool
	contadorConexiones int
	mu sync.Mutex
}
func newConexion(enchufe net.Conn, identificador int) *conexion{ 

	return &conexion{
		enchufe: enchufe,
		conexionActiva: true,
		identificador: identificador,
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
   	
	if c.lector.Scan(){
		return c.lector.Bytes(), nil
	}

	err := c.lector.Err()
	if err != nil{
		return nil, fmt.Errorf("error al leer mensaje: %w", err)
	}
	return nil, nil
}



func (c* conexion) enviaMensaje(builder *mensajeAClienteBuilder) error{
	c.mu.Lock()
    defer c.mu.Unlock()

	json, err := builder.construye()
	if err != nil{
		return fmt.Errorf("Error al construir el JSON: %w", err) 
	}

	jsonSalto := append([]byte(json), '\n')
	_, err = c.imprimir.Write(jsonSalto)
	if err != nil{
		return fmt.Errorf("Error al enviar por el socket: %w", err)
	}

	err = c.imprimir.Flush()
	if err != nil{
		return fmt.Errorf("Error en flush para el socket: %w", err) 
	}
	return nil
}


