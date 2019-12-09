package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

func createDummyAdoptableWithName(name string) *entities.Adoptable {
	return entities.CreateAdoptableWithName(name)
}

var _ = Describe("person", func() {
	defaultEmail := "dummy@tld.rd"

	It("should create an instance", func() {
		name := "Janie Soto"

		actual := entities.CreatePersonWithNameAndEmail(name, defaultEmail)

		Expect(actual).ToNot(BeNil())
		Expect(actual.Adoptables).To(HaveLen(0))
	})

	It("should create adopt 2 adoptables", func() {
		name := "Myra Wise"
		person := entities.CreatePersonWithNameAndEmail(name, defaultEmail)

		for i := 1; i <= 2; i++ {
			adoptableName := fmt.Sprintf("Dummy %d", i)
			adoptable := entities.CreateAdoptableWithName(adoptableName)
			person.AdoptAdoptable(adoptable)
		}

		Expect(person.Adoptables).To(HaveLen(2))
		Expect(person.Adoptables[0].Name).To(Equal("Dummy 1"))
		Expect(person.Adoptables[1].Name).To(Equal("Dummy 2"))
	})
})
