package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details/adapters/usecases/entities"
)

func createDummyPersonWithNameAndEmailAndToolWithName(personName string, personEmail string, toolName string) *entities.Person {
	tool := entities.CreateToolWithName(toolName)
	person := entities.CreatePersonWithNameAndEmail(personName, personEmail)
	person.AdoptTool(tool)
	return person
}

var _ = Describe("adoption", func() {
	defaultEmail := "dummy@tld.rd"

	It("should create an instance", func() {
		actual := entities.BuildAdoption()

		Expect(actual).ToNot(BeNil())
		Expect(actual.People).To(HaveLen(0))
	})

	It("should include 2 persons", func() {
		adoption := entities.BuildAdoption()

		for i := 1; i <= 2; i++ {
			personName := fmt.Sprintf("Dummy %d", i)
			person := entities.CreatePersonWithNameAndEmail(personName, defaultEmail)
			adoption.IncludePerson(*person)
		}

		Expect(adoption.People).To(HaveLen(2))
		Expect(adoption.People[0].Name).To(Equal("Dummy 1"))
		Expect(adoption.People[1].Name).To(Equal("Dummy 2"))
	})

	It("should get 0 adoption when nobody adopt the tool", func() {
		adoption := entities.BuildAdoption()
		person := createDummyPersonWithNameAndEmailAndToolWithName("Stanley Sherman", defaultEmail, "Tool 1")
		adoption.IncludePerson(*person)
		expectedToolName := "Tool 2"
		expectedTool := entities.CreateToolWithName(expectedToolName)

		actual, _ := adoption.CalculateForTool(*expectedTool)

		Expect(actual).To(Equal(0))
	})

	It("should get 25 adoption when 1 of 4 people adopt the tool", func() {
		adoption := entities.BuildAdoption()
		person1 := createDummyPersonWithNameAndEmailAndToolWithName("Stanley Sherman", defaultEmail, "Tool 1")
		adoption.IncludePerson(*person1)
		person2 := createDummyPersonWithNameAndEmailAndToolWithName("Marie Holloway", defaultEmail, "Tool 2")
		adoption.IncludePerson(*person2)
		person3 := createDummyPersonWithNameAndEmailAndToolWithName("Fanny Watson", defaultEmail, "Tool 2")
		adoption.IncludePerson(*person3)
		person4 := createDummyPersonWithNameAndEmailAndToolWithName("Winifred McKinney", defaultEmail, "Tool 2")
		adoption.IncludePerson(*person4)
		expectedToolName := "Tool 1"
		expectedTool := entities.CreateToolWithName(expectedToolName)

		actual, _ := adoption.CalculateForTool(*expectedTool)

		Expect(actual).To(Equal(25))
	})
})
