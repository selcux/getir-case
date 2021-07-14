package util_test

import (
	"getir-case/internal/model/fetch"
	"getir-case/internal/service/persistent"
	"time"
)

var (
	invalidRequestToDataQueryCases = []struct {
		name  string
		input *fetch.Request
	}{
		{
			name: "invalid StartDate format",
			input: &fetch.Request{
				StartDate: "23344wer",
				EndDate:   "2018-02-02",
				MinCount:  2000,
				MaxCount:  3000,
			},
		},
		{
			name: "invalid EndDate format",
			input: &fetch.Request{
				StartDate: "2016-01-26",
				EndDate:   "20120305",
				MinCount:  2000,
				MaxCount:  3000,
			},
		},
		{
			name: "invalid start date",
			input: &fetch.Request{
				StartDate: "2020-34-54",
				EndDate:   "2021-02-02",
				MinCount:  0,
				MaxCount:  0,
			},
		},
		{
			name: "invalid end date",
			input: &fetch.Request{
				StartDate: "2016-01-26",
				EndDate:   "2021-56-34",
				MinCount:  0,
				MaxCount:  0,
			},
		},
		{
			name: "invalid chronological date order",
			input: &fetch.Request{
				StartDate: "2021-02-02",
				EndDate:   "2018-02-02",
				MinCount:  0,
				MaxCount:  0,
			},
		},
	}

	validRequestToDataQueryCases = []struct {
		name     string
		input    *fetch.Request
		expected *persistent.DataQuery
	}{
		{
			name: "valid date formats",
			input: &fetch.Request{
				StartDate: "2016-01-26",
				EndDate:   "2018-02-02",
				MinCount:  0,
				MaxCount:  0,
			},
			expected: &persistent.DataQuery{
				StartDate: time.Date(2016, 01, 26, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2018, 02, 02, 0, 0, 0, 0, time.UTC),
				MinCount:  0,
				MaxCount:  0,
			},
		},
		{
			name: "consistent counts",
			input: &fetch.Request{
				StartDate: "2020-03-10",
				EndDate:   "2021-05-04",
				MinCount:  1234,
				MaxCount:  8765,
			},
			expected: &persistent.DataQuery{
				StartDate: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2021, 5, 4, 0, 0, 0, 0, time.UTC),
				MinCount:  1234,
				MaxCount:  8765,
			},
		},
	}

	validRecordsToResponsesCases = []struct {
		name     string
		input    []persistent.DataQueryRecord
		expected []fetch.RecordResponse
	}{
		{
			name: "unordered valid data",
			input: []persistent.DataQueryRecord{
				{
					Key:        "a",
					CreatedAt:  time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC),
					TotalCount: 1234,
				},
				{
					Key:        "b",
					CreatedAt:  time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC),
					TotalCount: 9876,
				},
			},
			expected: []fetch.RecordResponse{
				{
					Key:        "b",
					CreatedAt:  time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC),
					TotalCount: 9876,
				},
				{
					Key:        "a",
					CreatedAt:  time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC),
					TotalCount: 1234,
				},
			},
		},
	}
)
