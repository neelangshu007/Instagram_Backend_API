package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func WriteJson(w http.ResponseWriter, r *http.Request, v interface{}) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(jsonData) == 0 {
		http.Error(w, "No data found for the request", http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func ReadJson(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}
