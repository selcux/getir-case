package fetch

import (
	"encoding/json"
	"getir-case/internal/handler"
	"getir-case/internal/model/fetch"
	storage "getir-case/internal/service/persistent"
	"getir-case/internal/util"
	"github.com/go-playground/validator"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.post(w, r)
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	var request fetch.Request
	rb := handler.NewResponseBuilder(w)

	err := json.NewDecoder(r.Body).Decode(&request)
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
