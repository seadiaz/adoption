package entities_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/src/details/adapters/use_cases/entities"
)

func createDummyToolWithName(name string) *entities.Tool {
	return entities.CreateToolWithName(name)
}

var _ = Describe("person", func() {
	defaultEmail := "dummy@tld.rd"

	It("should create an instance", func() {
		name := "Janie Soto"

		actual := entities.CreatePersonWithNameAndEmail(name, defaultEmail)

		Expect(actual).ToNot(BeNil())
		Expect(actual.Tools).To(HaveLen(0))
	})

	It("should create adopt 2 tools", func() {
		name := "Myra Wise"
		person := entities.CreatePersonWithNameAndEmail(name, defaultEmail)

		for i := 1; i <= 2; i++ {
			toolName := fmt.Sprintf("Dummy %d", i)
			tool := entities.CreateToolWithName(toolName)
			person.AdoptTool(tool)
		}

		Expect(person.Tools).To(HaveLen(2))
		Expect(person.Tools[0].Name).To(Equal("Dummy 1"))
		Expect(person.Tools[1].Name).To(Equal("Dummy 2"))
	})
})
