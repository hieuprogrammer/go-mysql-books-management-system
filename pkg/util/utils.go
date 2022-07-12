package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, error := ioutil.ReadAll(request.Body); error == nil {
		if error := json.Unmarshal([]byte(body), x); error != nil {
			return
		}
	}
}
