package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details/adapters/use-cases/entities"
)

func createDummyToolWithName(name string) *Tool {
	return entities.BuildToolWithName(name)
}

var _ = Describe("person", func() {
	It("should create an instance", func() {
		name := "Janie Soto"

		actual := entities.BuildPerson(name)

		Expect(actual).ToNot(BeNil())
		Expect(actual.Tools).To(HaveLen(0))
	})

	It("should create adopt 2 tools", func() {
		name := "Myra Wise"
		person := entities.BuildPerson(name)

		for i := 1; i <= 2; i++ {
			toolName := fmt.Sprintf("Dummy %d", i)
			tool := entities.BuildToolWithName(toolName)
			person.AdoptTool(*tool)
		}

		Expect(person.Tools).To(HaveLen(2))
		Expect(person.Tools[0].Name).To(Equal("Dummy 1"))
		Expect(person.Tools[1].Name).To(Equal("Dummy 2"))
	})
})
