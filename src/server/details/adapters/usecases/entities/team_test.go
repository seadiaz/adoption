package entities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

var _ = Describe("team", func() {
	It("should create an instance", func() {
		actual := entities.CreateTeamWithName("Dummy Name")

		Expect(actual).ToNot(BeNil())
		Expect(actual.People).To(HaveLen(0))
	})

	It("should add a person to the team", func() {
		actual := entities.CreateTeamWithName("Dummy Name")
		person := entities.CreatePersonWithNameAndEmail("Edith Jenkins", "onzempu@gidesud.net")

		actual.AddPerson(person)

		Expect(actual.People).To(HaveLen(1))
		Expect(actual.People[0].Name).To(Equal("Edith Jenkins"))
	})

	It("should remove a person to the team", func() {
		actual := entities.CreateTeamWithName("Dummy Name")
		person1 := entities.CreatePersonWithNameAndEmail("Edith Jenkins", "onzempu@gidesud.net")
		person2 := entities.CreatePersonWithNameAndEmail("Laura Floyd", "teb@gigowhaw.pe")
		person3 := entities.CreatePersonWithNameAndEmail("Alvin Estrada", "givwiv@daz.si")
		person4 := entities.CreatePersonWithNameAndEmail("Nell Nguyen", "mat@cefarkub.bv")

		actual.AddPerson(person1)
		actual.AddPerson(person2)
		actual.AddPerson(person3)
		actual.AddPerson(person4)

		actual.RemovePerson(person1)
		actual.RemovePerson(person2)
		actual.RemovePerson(person4)

		Expect(actual.People).To(HaveLen(1))
		Expect(actual.People[0].Name).To(Equal("Alvin Estrada"))
	})
})
