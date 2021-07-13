package util

import (
	"getir-case/internal/model/fetch"
	persistent2 "getir-case/internal/service/persistent"
	"github.com/pkg/errors"
	"time"
)

func RequestToDataQuery(model fetch.Request) (*persistent2.DataQuery, error) {
	startDate, err := time.Parse("2006-01-02", model.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "invalid date format")
	}

	endDate, err := time.Parse("2006-01-02", model.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "invalid date format")
	}

	return &persistent2.DataQuery{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  model.MinCount,
		MaxCount:  model.MaxCount,
	}, nil
}

func RecordsToResponses(records []persistent2.DataQueryRecord) []fetch.RecordResponse {
	responses := make([]fetch.RecordResponse, 0)

	for _, record := range records {
		response := fetch.RecordResponse{
			Key:        record.Key,
			CreatedAt:  record.CreatedAt,
			TotalCount: record.TotalCount,
		}

		responses = append(responses, response)
	}

	return responses
}
