package inmemory

import (
	"encoding/json"
	"getir-case/internal/handler"
	"getir-case/internal/model/inmemory"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
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

	rb.JsonResponse(http.StatusOK, keyValue)
}

func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//key := r.URL.Query().Get("key")

	//jsonResponse(w, http.StatusOK, response)
}
