package handler

import (
	"encoding/json"
	"net/http"

	Restmodel "github.com/vinodborole/restapi-boilerplate/src/app/infra/rest/generated/server/go"
	u "github.com/vinodborole/restapi-boilerplate/src/app/utils"
)

//HandleErrorResponse handle error response
func HandleErrorResponse(w http.ResponseWriter, message string, code int32) {
	var errorResponse Restmodel.ModelError
	errorResponse.Message = message
	errorResponse.Code = code
	bytess, _ := json.Marshal(&errorResponse)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bytess)
}

//HandleGenericSuccess handle generic success response
func HandleGenericSuccess(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response = u.Message(true, message)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	u.Respond(w, response)
}

//HandleSuccessResponse handle success response
func HandleSuccessResponse(w http.ResponseWriter, domain interface{}) {
	bytess, _ := json.Marshal(domain)
	w.WriteHeader(http.StatusOK)
	w.Write(bytess)
}
