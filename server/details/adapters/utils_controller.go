package adapters

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/golang/glog"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

func replyWithError(w http.ResponseWriter, statusCode int, err error) {
	errResponse := &ErrorResponse{
		Message: err.Error(),
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResponse)
}

func replyJSONResponse(w http.ResponseWriter, output interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func validateRequest(r *http.Request, rules map[string]string, output interface{}) error {
	data, err := validate.FromRequest(r)
	if err != nil {
		glog.Error(err)
		return err
	}
	v := data.Create()
	v.ConfigRules(rules)
	if !v.Validate() {
		return errors.New(v.Errors.One())
	}
	v.BindSafeData(output)
	return nil
}

func getPathParam(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}

func getQueryParam(r *http.Request, name string) string {
	val := r.URL.Query().Get(name)
	return val
}

func getQueryParamMapKeyValue(r *http.Request, name string) (string, string) {
	queryVal := r.URL.Query().Get(name)
	pattern := regexp.MustCompile(`(\w+[\w \-_.]+):(\w+[\w \-_.]+)`)
	kv := pattern.FindStringSubmatch(queryVal)
	if len(kv) < 3 {
		return "", ""
	}

	return kv[1], kv[2]
}
