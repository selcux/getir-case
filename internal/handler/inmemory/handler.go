package inmemory

import (
	"encoding/json"
	"getir-case/internal/handler"
	"getir-case/internal/model/inmemory"
	storage "getir-case/internal/service/inmemory"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
	case http.MethodGet:
		h.get(w, r)
	}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	var keyValue inmemory.Request
	rb := handler.NewResponseBuilder(w)

	err := json.NewDecoder(r.Body).Decode(&keyValue)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	validate := validator.New()
	err = validate.Struct(keyValue)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	db, err := storage.NewStorage()
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	err = db.Set(keyValue.Key, keyValue.Value)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	rb.JsonResponse(http.StatusOK, keyValue)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	rb := handler.NewResponseBuilder(w)

	if key == "" {
		err := errors.New("key not given")
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
	}

	db, err := storage.NewStorage()
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	value, err := db.Get(key)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, inmemory.NewErrorResponse(err))
		return
	}

	response := &inmemory.Response{
		Key:   key,
		Value: value,
	}

	rb.JsonResponse(http.StatusOK, response)

}
