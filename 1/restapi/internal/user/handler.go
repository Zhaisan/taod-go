package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi/internal/handlers"
	"restapi/pkg/logging"
)

const (
	usersURL = "/users"
	userURL = "/users/:uuid"
)


//здесь будут логи и сервисы, всё что нужно хендлерам(get, update, etc)
type handler struct {
	logger *logging.Logger
}

//конструктор

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

//реализуем регистер который в хендлерах

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is the list of users"))  //пока заглушка
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create user "))  //пока заглушка
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is user by uuid"))  //пока заглушка
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))  //пока заглушка
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))  //пока заглушка
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))  //пока заглушка
}