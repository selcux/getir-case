package util_test

import (
	"github.com/joho/godotenv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "getir-case/testing"
)

func TestUtil(t *testing.T) {
	BeforeSuite(func() {
		err := godotenv.Load()
		Ω(err).ShouldNot(HaveOccurred())
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Util Suite")
}
