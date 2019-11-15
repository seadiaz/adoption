package details_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details"
)

func TestAll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "memory persistence")
}

var _ = Describe("create", func() {
	It("should add a single object", func() {
		persistence := details.BuildMemoryPersistence()
		id := "4128cbf6-b279-46b3-ae19-9f90ea190978"
		value := struct{ id string }{id}

		error := persistence.Create(id, value)

		actual, _ := persistence.Find(id)

		Expect(error).To(BeNil())
		Expect(actual).To(Equal(value))
	})

	It("should fail when id is empty", func() {
		persistence := details.BuildMemoryPersistence()
		id := ""
		value := struct{ id string }{id}

		error := persistence.Create(id, value)

		Expect(error).NotTo(BeNil())
	})
})

var _ = Describe("delete", func() {
	It("should remove one item", func() {
		persistence := details.BuildMemoryPersistence()
		id := "7879d950-e511-4798-a074-a951d9eddbb8"
		value := struct{ id string }{id: id}
		persistence.Create(id, value)

		error := persistence.Delete(id)
		actual, _ := persistence.Find(id)

		Expect(error).To(BeNil())
		Expect(actual).To(BeNil())

	})

	It("should fail when id is empty", func() {
		persistence := details.BuildMemoryPersistence()
		id := ""

		error := persistence.Delete(id)

		Expect(error).NotTo(BeNil())
	})
})

var _ = Describe("get all", func() {
	It("should return all values", func() {
		persistence := details.BuildMemoryPersistence()
		id := "7879d950-e511-4798-a074-a951d9eddbb8"
		value := struct{ id string }{id: id}
		persistence.Create(id, value)
		expected := make([]interface{}, 1)
		expected[0] = value

		results := persistence.GetAll()

		Expect(results).To(Equal(expected))
	})
})
