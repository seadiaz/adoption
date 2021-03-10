package docs

import "github.com/seadiaz/adoption/server/details/adapters"

// swagger:route GET /people people getPeople
// Fetch the all the people
// responses:
//   200: personResponse

// swagger:route POST /people people postPeople
// Create a new person
// responses:
//   200: personResponse

// This text will appear as description of your response body.
// swagger:response personResponse
type personResponseWrapper struct {
	// in:body
	Body adapters.PersonResponse
}

// swagger:parameters postPeople
type personRequestWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body adapters.PersonRequest
}
