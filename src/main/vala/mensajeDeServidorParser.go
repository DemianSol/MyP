package main

import {
	"encoding/json"
	"fmt"
}



// gobyexample.com/interfaces
type Mensaje interface {
	TipoMensaje() string
}

// https://gobyexample.com/json
type mensajeBasico struct{
	Type string `json:"type"`
}

type mensajeNewUser struct{
	Type string `json:"type"`
	Username string `json:"username"`
}

func (m mensajeNewUser) TipoMensaje() string{
	return m.Type
}

type mensajeNewStatus struct{
	Type string `json:"type"`
	Username string `json:"username"`
	Status string `json:"status"`
}

func (m mensajeNewUser) TipoMensaje() string{
	return m.Type
}

type mensajeListaUsuarios struct{
	Type string `json:"type"`
	Username map[string]string `json:"users"`
}

func (m mensajeListaUsuarios) TipoMensaje() string{
	return m.Type
}

type mensajeTextoDe struct{
	Type string `json:"type"`
	Username string `json:"username"`
	Text string `json:"text"`
}

func (m mensajeTextoDe) TipoMensaje() string{
	return m.Type
}

type mensajeTextoPublico struct{
	Type string `json:"type"`
	Username string `json:"username"`
	Text string `json:"text"`
}

func (m mensajeTextoPublico) TipoMensaje() string{
	return m.Type
}

type mensajeUnidoASala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`
}

func (m mensajeUnidoASala) TipoMensaje() string{
	return m.Type
}

type mensajeUsuariosSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Users map[string]string `json:"users"`
}

func (m mensajeUsuariosSala) TipoMensaje() string{
	return m.Type
}

type mensajeTextoSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`
	Text string `json:"text"`
}

func (m mensajeTextoSala) TipoMensaje() string{
	return m.Type
}

type mensajeDejaSala struct{
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`
}

func (m mensajeDejaSala) TipoMensaje() string{
	return m.Type
}

type mensajeDesconectado struct{
	Type string `json:"type"`
	Username string `json:"username"`
}

func (m mensajeDesconectado) TipoMensaje() string{
	return m.Type
}

// https://go.dev/blog/json
// https://go.dev/doc/effective_go#composite_literals
//https://go.dev/doc/effective_go#allocation_new
// https://stackoverflow.com/questions/18125625/constructors-in-go   idea para añadir valor inicial a map
func ProcesaJSON(json []byte) (Mensaje, error){
	var mensaje mensajeBasico
	if err := json.Unmarshal(json, &mensaje); err != nil{
		return nil, fmt.Errorf("Error en el formato del mensaje:", err)
	}

	switch mensaje.Type {
	case "NEW_USER":
		var msj mensajeNewUser
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil
	
	case "NEW_STATUS":
		var msj mensajeNewStatus
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "USER_LIST":
		var msj mensajeListaUsuarios
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "TEXT_FROM":
		var msj mensajeTextoDe
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "PUBLIC_TEXT_FROM":
		var msj mensajeTextoPublico
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "JOINED_ROOM":
		var msj mensajeUnidoASala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "ROOM_USER_LIST":
		var msj mensajeUsuariosSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "ROOM_TEXT_FROM":
		var msj mensajeTextoDe
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "LEFT_ROOM":
		var msj mensajeDejaSala
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	case "DISCONNECTED":
		var msj mensajeDesconectado
		if err := json.Unmarshal(json, &msj); err != nil{
			return nil, err
		}
		return msj, nil

	default:
		return nil, fmt.Errorf("El mensaje no se reconoce como un tipo válido:", mensaje.Type)
	}
}

