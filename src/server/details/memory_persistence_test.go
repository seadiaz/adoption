package details_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/server/details"
	"github.com/seadiaz/adoption/src/server/details/adapters"
)

const persistenceTypeDummy = "people"

type dummyPersistedData struct {
	ID string
}

func (t *dummyPersistedData) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *dummyPersistedData) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}

// Clone ...
func (t *dummyPersistedData) Clone() adapters.PersistedData {
	return &dummyPersistedData{}
}

func TestAll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "memory persistence")
}

var _ = Describe("create", func() {
	It("should add a single object", func() {
		persistence := details.BuildMemoryPersistence()
		id := "4128cbf6-b279-46b3-ae19-9f90ea190978"
		// value := struct{ id string }{id}
		value := &dummyPersistedData{id}

		error := persistence.Create(persistenceTypeDummy, id, value)

		actual, _ := persistence.Find(persistenceTypeDummy, id, value)

		Expect(error).To(BeNil())
		Expect(actual).To(Equal(value))
	})

	It("should fail when id is empty", func() {
		persistence := details.BuildMemoryPersistence()
		id := ""
		value := &dummyPersistedData{id}

		error := persistence.Create(persistenceTypeDummy, id, value)

		Expect(error).NotTo(BeNil())
	})
})

var _ = Describe("delete", func() {
	It("should remove one item", func() {
		persistence := details.BuildMemoryPersistence()
		id := "7879d950-e511-4798-a074-a951d9eddbb8"
		value := &dummyPersistedData{id}
		persistence.Create(persistenceTypeDummy, id, value)

		error := persistence.Delete(persistenceTypeDummy, id)
		actual, _ := persistence.Find(persistenceTypeDummy, id, value)

		Expect(error).To(BeNil())
		Expect(actual).To(BeNil())

	})

	It("should fail when id is empty", func() {
		persistence := details.BuildMemoryPersistence()
		id := ""

		error := persistence.Delete(persistenceTypeDummy, id)

		Expect(error).NotTo(BeNil())
	})
})

var _ = Describe("get all", func() {
	It("should return all values", func() {
		persistence := details.BuildMemoryPersistence()
		id := "7879d950-e511-4798-a074-a951d9eddbb8"
		value := &dummyPersistedData{id}
		persistence.Create(persistenceTypeDummy, id, value)
		expected := make([]interface{}, 1)
		expected[0] = value

		results, _ := persistence.GetAll(persistenceTypeDummy, value)

		Expect(results).To(Equal(expected))
	})
})
