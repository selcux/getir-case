package util_test

import (
	"fmt"
	"getir-case/internal/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert", func() {
	Describe("Convert Request to DataQuery", func() {
		Context("With an invalid date format", func() {
			for _, test := range invalidRequestToDataQueryCases {
				Describe(test.name, func() {
					It("should be an error", func() {
						actual, err := util.RequestToDataQuery(test.input)
						Ω(err).Should(HaveOccurred())
						Ω(actual).Should(BeNil())
					})
				})
			}
		})

		Context("With a valid data", func() {
			for _, test := range validRequestToDataQueryCases {
				It(fmt.Sprintf("should be %s", test.name), func() {
					actual, err := util.RequestToDataQuery(test.input)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(actual).Should(Equal(test.expected))
				})
			}
		})
	})

	Describe("Convert []DataQueryRecord to []RecordResponse", func() {
		Context("With a valid (un)ordered array", func() {
			for _, test := range validRecordsToResponsesCases {
				It(test.name, func() {
					actual := util.RecordsToResponses(test.input)
					Ω(len(actual)).Should(Equal(len(test.expected)))
					Ω(actual).Should(ContainElements(test.expected))
				})
			}
		})
	})
})
