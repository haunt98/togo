package transports

import (
	"encoding/json"
	"net/http"
)

// Make JSON response with data and error (if exist)
func makeJSONResponse(rsp http.ResponseWriter, statusCode int, data interface{}, err error) {
	rsp.Header().Set("Content-Type", "application/json")
	rsp.WriteHeader(statusCode)

	if err != nil {
		json.NewEncoder(rsp).Encode(map[string]interface{}{
			errorField: err.Error(),
		})
		return
	}

	json.NewEncoder(rsp).Encode(map[string]interface{}{
		dataField: data,
	})
}
