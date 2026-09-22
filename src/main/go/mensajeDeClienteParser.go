package main

import {
	"encoding/json"
	"fmt"
}



// gobyexample.com/interfaces
type mensaje interface {
	tipoMensaje() string
}

// https://gobyexample.com/json
type mensajeBasico struct{
	Type string `json:"type"`
}

type mensajeIdentificar struct{
	Type string `json:"type"`
	Username string `json:"username"`
}

func (m mensajeIdentificar) tipoMensaje() string{
	return m.Type
}

type mensajeNewStatus struct{
	Type string `json:"type"`
	Status string `json:"status"`
}

func (m mensajeNewStatus) tipoMensaje() string{
	return m.Type
}

type mensajeListaUsuarios struct{
	Type string `json:"type"`
}

func (m mensajeListaUsuarios) tipoMensaje() string{
	return m.Type
}

type mensajeTexto struct{
	Type string `json:"type"`
	Username string `json:"username"`
	Text string `json:"text"`
}

func (m mensajeTexto) tipoMensaje() string{
	return m.Type
}

type mensajeTextoPublico struct{
	Type string `json:"type"`
	Username string `json:"username"`
	Text string `json:"text"`
}

func (m mensajeTextoPublico) tipoMensaje() string{
	return m.Type
}

type mensajeNuevaSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (m mensajeNuevaSala) tipoMensaje() string{
	return m.Type
}

type mensajeInvita struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Users map[string]string `json:"usernames"`
}

func (m mensajeInvita) tipoMensaje() string{
	return m.Type
}

type mensajeUnirSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (m mensajeUnirSala) tipoMensaje() string{
	return m.Type
}

type mensajeUsuariosSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (m mensajeUsuariosSala) tipoMensaje() string{
	return m.Type
}

type mensajeTextoSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Text string `json:"text"`
}

func (m mensajeTextoSala) tipoMensaje() string{
	return m.Type
}


type mensajeDejaSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (m mensajeDejaSala) tipoMensaje() string{
	return m.Type
}

type mensajeDesconectado struct{
	Type string `json:"type"`
}

func (m mensajeDesconectado) tipoMensaje() string{
	return m.Type
}

// https://go.dev/blog/json
// https://go.dev/doc/effective_go#composite_literals
//https://go.dev/doc/effective_go#allocation_new
// https://stackoverflow.com/questions/18125625/constructors-in-go   idea para añadir valor inicial a map
func procesaJSON(json []byte) (mensaje, error){
	var mensaje mensajeBasico
	if err := json.Unmarshal(json, &mensaje); err != nil{
		return nil, fmt.Errorf("Error en el formato del mensaje:", err)
	}

	switch mensaje.Type {
	case "IDENTIFY":
		var msj mensajeIdentificar
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil
	
	case "STATUS":
		var msj mensajeNewStatus
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "USERS":
		var msj mensajeListaUsuarios
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "TEXT":
		var msj mensajeTexto
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "PUBLIC_TEXT":
		var msj mensajeTextoPublico
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "NEW_ROOM":
		var msj mensajeNuevaSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "INVITE":
		var msj mensajeInvita
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "JOIN_ROOM":
		var msj mensajeUnirSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "ROOM_USERS":
		var msj mensajeUsuariosSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "ROOM_TEXT":
		var msj mensajeTextoSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil
	
	case "LEAVE_ROOM":
		var msj mensajeDejaSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil
	
	case "DISCONNECT":
		var msj mensajeDesconectado
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	default:
		return nil, fmt.Errorf("El mensaje no se reconoce como un tipo válido:", mensaje.Type)
	}
}

