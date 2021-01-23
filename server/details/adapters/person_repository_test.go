package adapters_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/server/details"
	"github.com/seadiaz/adoption/server/details/adapters"
	"github.com/seadiaz/adoption/server/details/adapters/usecases/entities"
)

func TestAll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "person repository")
}

var _ = Describe("save", func() {
	It("should not save duplicate email", func() {
		persistence := details.BuildMemoryPersistence()
		repository := adapters.CreatePersonRepository(persistence)
		person := entities.CreatePersonWithNameAndEmail("Dummy 1", "dummy@tld.com")
		person2 := entities.CreatePersonWithNameAndEmail("Dummy 2", "dummy@tld.com")
		repository.SavePerson(person)
		repository.SavePerson(person2)

		people, _ := repository.GetAllPeople()

		Expect(people).To(HaveLen(1))
	})
})
