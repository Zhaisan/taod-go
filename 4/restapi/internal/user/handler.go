package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi/internal/apperror"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut,userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this is API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test", "test", "t123")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))  //пока заглушка

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))  //пока заглушка

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))  //пока заглушка

	return nil
}