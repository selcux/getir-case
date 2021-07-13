package inmemory

import (
	"encoding/json"
	"getir-case/internal/handler"
	"getir-case/internal/model/inmemory"
	storage "getir-case/internal/service/inmemory"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
