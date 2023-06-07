package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReturnJson(structModel any) []byte {

	bytes, _ := json.Marshal(structModel)

	return bytes

}

func ParseBody(r *http.Request, x interface{}) {
	if body, error := ioutil.ReadAll(r.Body); error == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}

func ReturnResponse(w http.ResponseWriter, returnData any, statusCode int) {

	var jsonObj = ReturnJson(returnData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonObj)
}
