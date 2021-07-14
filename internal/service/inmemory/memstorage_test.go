package inmemory_test

import (
	"getir-case/internal/service/inmemory"
	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

func newTestRedis() *redismock.ClientMock {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}

var _ = Describe("Memstorage", func() {
	var rClient *miniredis.Miniredis

	BeforeEach(func() {
		s, err := miniredis.Run()
		Ω(err).ShouldNot(HaveOccurred())

		rClient = s
		os.Setenv("REDIS_URL", "redis://"+rClient.Addr())
	})

	AfterEach(func() {
		os.Unsetenv("REDIS_URL")
		rClient.Close()
	})

	It("should be able to connect", func() {
		_, err := inmemory.NewStorage()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should set a key value", func() {
		storage, err := inmemory.NewStorage()
		Ω(err).ShouldNot(HaveOccurred())

		err = storage.Set("a", "x")
		Ω(err).ShouldNot(HaveOccurred())

		Ω(rClient.Exists("a")).Should(BeTrue())
	})

	It("should get the set value", func() {
		storage, err := inmemory.NewStorage()
		Ω(err).ShouldNot(HaveOccurred())

		err = storage.Set("a", "x")
		Ω(err).ShouldNot(HaveOccurred())

		actual, err := storage.Get("a")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(actual).Should(Equal("x"))

		stored, err := rClient.Get("a")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(stored).Should(Equal("x"))
	})
})
