package main

import (
	"encoding/json"
	"fmt"
)


const(
	usuarioNuevo = "NEW_USER"
	nuevoStatus = "NEW_STATUS"
	listaUsuario = "USER_LIST"
	textoDesde = "TEXTO_FROM"
	textoPublico = "PUBLIC_TEXT_FROM"
	unidoASala = "JOINED_ROOM"
	usuariosSala = "ROOM_USER_LIST"
	textoSala = "ROOM_TEXT_FROM"
	dejaSala = "LEFT_ROOM"
	desconectado = "DISCONNECTED"
	respuesta = "RESPONSE"
	invitacion = "INVITATION"
	desconectar = "DISCONNECT"
	texto = "TEXT" // esto lo agergue
)

// https://gobyexample.com/json
type mensajeAClienteBuilder struct{
	data map[string]any
}

// https://go.dev/doc/effective_go#composite_literals
//https://go.dev/doc/effective_go#allocation_new
// https://stackoverflow.com/questions/18125625/constructors-in-go   idea para añadir valor inicial a map
func newMensajeAClienteBuilder(tipoMensaje string) *mensajeAClienteBuilder {
	m := new(mensajeAClienteBuilder)
	m.data = make(map[string]any)
	m.data["type"] = tipoMensaje
	return m
}

func (m *mensajeAClienteBuilder) añadir(llave string, valor string) *mensajeAClienteBuilder{
	m.data[llave] = valor
	return m
}

func (m *mensajeAClienteBuilder) añadeListaUsuarios(usuarios map[string]string) *mensajeAClienteBuilder{
	m.data["users"] = usuarios
	return m
}

func (m *mensajeAClienteBuilder) añadirRespuesta(operacion string, resultado string, extra string)*mensajeAClienteBuilder{
	m.data["type"] = respuesta
	m.data["operation"] = operacion
	m.data["result"] = resultado
	if (extra != ""){
		m.data["extra"] = extra
	}
	return m

}

func (m *mensajeAClienteBuilder) construye() (string, error){
	j, err := json.Marshal(m.data)
	if err != nil{
		return "", fmt.Errorf("Error al construir el JSON para el cliente:%w", err)		
	}
	return string(j) + "\n", nil
}