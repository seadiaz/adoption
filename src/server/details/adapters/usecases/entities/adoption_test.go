package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

func createDummyPersonWithNameAndEmailAndAdoptable(personName string, personEmail string, tool *entities.Adoptable) *entities.Person {
	person := entities.CreatePersonWithNameAndEmail(personName, personEmail)
	person.AdoptAdoptable(tool)
	return person
}

var _ = Describe("adoption", func() {
	defaultEmail := "dummy@tld.rd"

	It("should create an instance", func() {
		actual := entities.CreateAdoption()

		Expect(actual).ToNot(BeNil())
		Expect(actual.People).To(HaveLen(0))
	})

	It("should include 2 persons", func() {
		adoption := entities.CreateAdoption()

		for i := 1; i <= 2; i++ {
			personName := fmt.Sprintf("Dummy %d", i)
			person := entities.CreatePersonWithNameAndEmail(personName, defaultEmail)
			adoption.IncludePerson(person)
		}

		Expect(adoption.People).To(HaveLen(2))
		Expect(adoption.People[0].Name).To(Equal("Dummy 1"))
		Expect(adoption.People[1].Name).To(Equal("Dummy 2"))
	})

	It("should get 0 adoption when nobody adopt the tool", func() {
		adoption := entities.CreateAdoption()
		tool := createDummyAdoptableWithName("Adoptable 1")
		person := createDummyPersonWithNameAndEmailAndAdoptable("Stanley Sherman", defaultEmail, tool)
		adoption.IncludePerson(person)
		expectedAdoptableName := "Adoptable 2"
		expectedAdoptable := entities.CreateAdoptableWithName(expectedAdoptableName)

		actual := adoption.CalculateForAdoptable(expectedAdoptable)

		Expect(0).To(Equal(actual))
	})

	It("should get 25 adoption when 1 of 4 people adopt the tool", func() {
		adoption := entities.CreateAdoption()
		tool1 := createDummyAdoptableWithName("Adoptable 1")
		tool2 := createDummyAdoptableWithName("Adoptable 2")
		person1 := createDummyPersonWithNameAndEmailAndAdoptable("Stanley Sherman", defaultEmail, tool1)
		adoption.IncludePerson(person1)
		person2 := createDummyPersonWithNameAndEmailAndAdoptable("Marie Holloway", defaultEmail, tool2)
		adoption.IncludePerson(person2)
		person3 := createDummyPersonWithNameAndEmailAndAdoptable("Fanny Watson", defaultEmail, tool2)
		adoption.IncludePerson(person3)
		person4 := createDummyPersonWithNameAndEmailAndAdoptable("Winifred McKinney", defaultEmail, tool2)
		adoption.IncludePerson(person4)

		expectedAdoptable := tool1

		actual := adoption.CalculateForAdoptable(expectedAdoptable)

		Expect(25).To(Equal(actual))
	})
})
