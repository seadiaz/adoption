package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details/adapters/use-cases/entities"
)

func createDummyPersonWithNameAndToolName(personName string, toolName string) *entities.Person {
	tool := entities.BuildToolWithName(toolName)
	person := entities.CreatePersonWithName(personName)
	person.AdoptTool(tool)
	return person
}

var _ = Describe("adoption", func() {
	It("should create an instance", func() {
		actual := entities.BuildAdoption()

		Expect(actual).ToNot(BeNil())
		Expect(actual.People).To(HaveLen(0))
	})

	It("should include 2 persons", func() {
		adoption := entities.BuildAdoption()

		for i := 1; i <= 2; i++ {
			personName := fmt.Sprintf("Dummy %d", i)
			person := entities.CreatePersonWithName(personName)
			adoption.IncludePerson(*person)
		}

		Expect(adoption.People).To(HaveLen(2))
		Expect(adoption.People[0].Name).To(Equal("Dummy 1"))
		Expect(adoption.People[1].Name).To(Equal("Dummy 2"))
	})

	It("should get 0 adoption when nobody adopt the tool", func() {
		adoption := entities.BuildAdoption()
		person := createDummyPersonWithNameAndToolName("Stanley Sherman", "Tool 1")
		adoption.IncludePerson(*person)
		expectedToolName := "Tool 2"
		expectedTool := entities.BuildToolWithName(expectedToolName)

		actual, _ := adoption.CalculateForTool(*expectedTool)

		Expect(actual).To(Equal(0))
	})

	It("should get 25 adoption when 1 of 4 people adopt the tool", func() {
		adoption := entities.BuildAdoption()
		person1 := createDummyPersonWithNameAndToolName("Stanley Sherman", "Tool 1")
		adoption.IncludePerson(*person1)
		person2 := createDummyPersonWithNameAndToolName("Marie Holloway", "Tool 2")
		adoption.IncludePerson(*person2)
		person3 := createDummyPersonWithNameAndToolName("Fanny Watson", "Tool 2")
		adoption.IncludePerson(*person3)
		person4 := createDummyPersonWithNameAndToolName("Winifred McKinney", "Tool 2")
		adoption.IncludePerson(*person4)
		expectedToolName := "Tool 1"
		expectedTool := entities.BuildToolWithName(expectedToolName)

		actual, _ := adoption.CalculateForTool(*expectedTool)

		Expect(actual).To(Equal(25))
	})
})
