package author

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi/internal/apperror"
	"restapi/internal/handlers"
	"restapi/pkg/logging"
)

const (
	authorsURL = "/authors"
	authorURL = "/authors/:uuid"
)


//здесь будут логи и сервисы, всё что нужно хендлерам(get, update, etc)
type handler struct {
	logger *logging.Logger
	repository Repository
}

//конструктор

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger: logger,
	}
}

//реализуем регистер который в хендлерах

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, apperror.Middleware(h.GetList))
	//router.HandlerFunc(http.MethodPost, authorsURL, apperror.Middleware(h.CreateAuthor))
	//router.HandlerFunc(http.MethodGet, authorURL, apperror.Middleware(h.GetAuthorByUUID))
	//router.HandlerFunc(http.MethodPut, authorURL, apperror.Middleware(h.UpdateAuthor))
	//router.HandlerFunc(http.MethodPatch, authorURL, apperror.Middleware(h.PartiallyUpdateAuthor))
	//router.HandlerFunc(http.MethodDelete, authorURL, apperror.Middleware(h.DeleteAuthor))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

//func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) error {
//	return fmt.Errorf("this is API error")
//}
//
//func (h *handler) GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
//	return apperror.NewAppError(nil, "test", "test", "t123")
//}
//
//func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
//	w.WriteHeader(204)
//	w.Write([]byte("this is update author"))  //пока заглушка
//
//	return nil
//}
//
//func (h *handler) PartiallyUpdateAuthor(w http.ResponseWriter, r *http.Request) error{
//	w.WriteHeader(204)
//	w.Write([]byte("this is partially update author"))  //пока заглушка
//
//	return nil
//}
//
//func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
//	w.WriteHeader(204)
//	w.Write([]byte("this is delete author"))  //пока заглушка
//
//	return nil
//}
