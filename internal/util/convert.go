package util

import (
	"getir-case/internal/db"
	"getir-case/internal/model/fetch"
	"github.com/pkg/errors"
	"time"
)

func RequestToDataQuery(model fetch.Request) (*db.DataQuery, error) {
	startDate, err := time.Parse("2006-01-02", model.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "invalid date format")
	}

	endDate, err := time.Parse("2006-01-02", model.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "invalid date format")
	}

	return &db.DataQuery{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  model.MinCount,
		MaxCount:  model.MaxCount,
	}, nil
}

func RecordsToResponses(records []db.DataQueryRecord) []fetch.RecordResponse {
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
