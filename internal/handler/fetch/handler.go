package fetch

import (
	"encoding/json"
	"getir-case/internal/handler"
	"getir-case/internal/model"
	"getir-case/internal/model/fetch"
	storage "getir-case/internal/service/persistent"
	"getir-case/internal/util"
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
	rb := handler.NewResponseBuilder(w)

	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
	default:
		err := errors.New("not found")
		rb.JsonResponse(http.StatusNotFound, model.NewErrorResponse(err))
	}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	request := new(fetch.Request)
	rb := handler.NewResponseBuilder(w)

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, fetch.NewErrorResponse(err))
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, fetch.NewErrorResponse(err))
		return
	}

	database, err := storage.NewDb()
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, fetch.NewErrorResponse(err))
		return
	}
	defer database.Disconnect()

	query, err := util.RequestToDataQuery(request)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, fetch.NewErrorResponse(err))
		return
	}

	queryRecords, err := database.Get(query)
	if err != nil {
		rb.JsonResponse(http.StatusBadRequest, fetch.NewErrorResponse(err))
		return
	}

	responseRecords := util.RecordsToResponses(queryRecords)
	response := &fetch.Response{
		Code:    0,
		Msg:     "Success",
		Records: responseRecords,
	}

	rb.JsonResponse(http.StatusOK, response)
}
