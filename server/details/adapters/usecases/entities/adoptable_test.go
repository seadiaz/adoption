package entities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seadiaz/adoption/server/details/adapters/usecases/entities"
)

var _ = Describe("adoptable", func() {
	It("should create an instance", func() {
		actual := entities.CreateAdoptableWithNameAndStrategy("Dummy Name", entities.StrategyTypeSingle)

		Expect(actual).ToNot(BeNil())
		Expect(actual.Labels).To(HaveLen(0))
	})

	It("should add a label to the adoptable", func() {
		actual := entities.CreateAdoptableWithNameAndStrategy("Dummy Name", entities.StrategyTypeSingle)
		label := entities.CreateLabelWithKindAndValue("tag", "my-tag")

		actual.AddLabel(label)

		Expect(actual.Labels).To(HaveLen(1))
		Expect(actual.Labels[0].Kind).To(Equal("tag"))
		Expect(actual.Labels[0].Value).To(Equal("my-tag"))
	})

	It("should replace a label to the adoptable", func() {
		actual := entities.CreateAdoptableWithNameAndStrategy("Dummy Name", entities.StrategyTypeSingle)
		label := entities.CreateLabelWithKindAndValue("tag", "my-tag")
		label2 := entities.CreateLabelWithKindAndValue("tag", "my-new-tag")

		actual.AddLabel(label)
		actual.AddLabel(label2)

		Expect(actual.Labels).To(HaveLen(1))
		Expect(actual.Labels[0].Kind).To(Equal("tag"))
		Expect(actual.Labels[0].Value).To(Equal("my-new-tag"))
	})

	It("should remove a label to the adoptable", func() {
		actual := entities.CreateAdoptableWithNameAndStrategy("Dummy Name", entities.StrategyTypeSingle)
		label := entities.CreateLabelWithKindAndValue("tag", "my-tag")
		actual.AddLabel(label)

		actual.RemoveLabel(label)

		Expect(actual.Labels).To(HaveLen(0))
	})
})
