package fetch

import (
	"encoding/json"
	"getir-case/internal/db"
	"getir-case/internal/handler"
	"getir-case/internal/model/fetch"
	"getir-case/internal/util"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	database, err := db.NewDb()
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
