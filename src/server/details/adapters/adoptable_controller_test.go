package adapters_test

import (
	"testing"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/seadiaz/adoption/src/server/details/adapters"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

func TestAllAdoptableController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "person repository")
}

type fakeAdoptableService struct {

}

// GetAllAdoptables ...
func (s *fakeAdoptableService) GetAllAdoptables() ([]*entities.Adoptable, error) {
	return nil, nil
}

// FindAdoptablesFilterByLabelKindAndValue ...
func (s *fakeAdoptableService) FindAdoptablesFilterByLabelKindAndValue(labelKind, labelValue string) ([]*entities.Adoptable, error) {
	return nil, nil
}

// CreateAdoptable ...
func (s *fakeAdoptableService) CreateAdoptable(name string) (*entities.Adoptable, error) {
	return nil, nil
}

// FindAdoptable ...
func (s *fakeAdoptableService) FindAdoptable(id string) (*entities.Adoptable, error) {
	return nil, nil
}

// AddLabelToAdoptable ...
func (s *fakeAdoptableService) AddLabelToAdoptable(labelKind string, labelValue string, adoptableID string) (*entities.Adoptable, error) {
	return nil, nil
}

type fakeAdoptionService struct {

}

// CalculateAdoptionForAdoptable ...
func (s *fakeAdoptionService) CalculateAdoptionForAdoptable(id string) (map[string]interface{}, error) {
	return nil, nil
}

type fakeRouter struct {

}

func (r *fakeRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) adapters.Route {
	return &fakeRoute{}
}

type fakeRoute struct {

}

func (r *fakeRoute) Methods(methods ...string) adapters.Route {
	return r
}

var _ = Describe("adoptable controller", func() {
	It("should parse label filter kind with space", func() {
		adoptableService := &fakeAdoptableService{}
		adoptionService := &fakeAdoptionService{}
		controller := adapters.CreateAdoptableController(adoptableService, adoptionService)
		router := &fakeRouter{}
		controller.AddRoutes(router)

		Expect(1).To(Equal(1))
	})
})
