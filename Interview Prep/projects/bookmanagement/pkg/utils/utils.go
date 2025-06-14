package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal([]byte(body), x); err == nil {
			fmt.Println(x)
			return
		} else {
			fmt.Println(err)
		}
	}
	// return nil, errors.New("Something went wrong")
}
