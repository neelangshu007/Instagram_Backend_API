package functions

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetLimitAndOffset(w http.ResponseWriter, r *http.Request) (int64, int64, error) {
	query := r.URL.Query()
	str := query.Get("limit")
	if len(str) == 0 {
		return 0, 0, fmt.Errorf("limit is not given in query parameter")
	}
	limit, err := strconv.Atoi(str)
	if err != nil {
		return 0, 0, fmt.Errorf("limit needs to be an integer")
	}
	str = query.Get("offset")
	if len(str) == 0 {
		return 0, 0, fmt.Errorf("offset is not given in query parameter")
	}
	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, 0, fmt.Errorf("offset needs to be an integer")
	}

	return int64(limit), int64(offset), nil
}
