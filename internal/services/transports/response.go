package transports

import (
	"encoding/json"
	"log"
	"net/http"
)

// Make JSON response with data and error (if exist)
func makeJSONResponse(rsp http.ResponseWriter, statusCode int, data interface{}, err error) {
	rsp.Header().Set("Content-Type", "application/json")
	rsp.WriteHeader(statusCode)

	if err != nil {
		if jsonErr := json.NewEncoder(rsp).Encode(map[string]interface{}{
			errorField: err.Error(),
		}); jsonErr != nil {
			log.Printf("failed to encode json: %s", jsonErr)
		}
		return
	}

	if jsonErr := json.NewEncoder(rsp).Encode(map[string]interface{}{
		dataField: data,
	}); jsonErr != nil {
		log.Printf("failed to encode json: %s", jsonErr)
	}
}
