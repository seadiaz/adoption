package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

func createDummyPersonWithNameAndEmailAndTool(personName string, personEmail string, tool *entities.Tool) *entities.Person {
	person := entities.CreatePersonWithNameAndEmail(personName, personEmail)
	person.AdoptTool(tool)
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
		tool := createDummyToolWithName("Tool 1")
		person := createDummyPersonWithNameAndEmailAndTool("Stanley Sherman", defaultEmail, tool)
		adoption.IncludePerson(person)
		expectedToolName := "Tool 2"
		expectedTool := entities.CreateToolWithName(expectedToolName)

		actual := adoption.CalculateForTool(expectedTool)

		Expect(0).To(Equal(actual))
	})

	It("should get 25 adoption when 1 of 4 people adopt the tool", func() {
		adoption := entities.CreateAdoption()
		tool1 := createDummyToolWithName("Tool 1")
		tool2 := createDummyToolWithName("Tool 2")
		person1 := createDummyPersonWithNameAndEmailAndTool("Stanley Sherman", defaultEmail, tool1)
		adoption.IncludePerson(person1)
		person2 := createDummyPersonWithNameAndEmailAndTool("Marie Holloway", defaultEmail, tool2)
		adoption.IncludePerson(person2)
		person3 := createDummyPersonWithNameAndEmailAndTool("Fanny Watson", defaultEmail, tool2)
		adoption.IncludePerson(person3)
		person4 := createDummyPersonWithNameAndEmailAndTool("Winifred McKinney", defaultEmail, tool2)
		adoption.IncludePerson(person4)

		expectedTool := tool1

		actual := adoption.CalculateForTool(expectedTool)

		Expect(25).To(Equal(actual))
	})
})
