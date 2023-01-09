package handlers

import "github.com/julienschmidt/httprouter"

//все сущности будут иметь хендлеры и все эти хендлеры будут реализовывать этот метод

type Handler interface {
	Register(router *httprouter.Router)
}
