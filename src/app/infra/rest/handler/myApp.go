package handler

import (
	"app/infra"
	"app/infra/rest/converter"
	"fmt"
	"net/http"
)

//About get app info
func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	success := true
	statusMsg := ""
	alog, _ := LogInit(r, "Info About my app")
	defer alog.LogMessageEnd(&success, &statusMsg)
	alog.LogMessageReceived()
	dbAppInfo, err := infra.GetUseCaseInteractor().Db.GetApp("restapi-boilerplate")
	if err != nil {
		success = false
		statusMsg = fmt.Sprintf("Error in getting app info, reason %s", err.Error())
		HandleErrorResponse(w, statusMsg, http.StatusNotFound)
		return
	}
	fitnessApp, _ := converter.GetAppSwaggerResponse(&dbAppInfo)
	HandleSuccessResponse(w, fitnessApp)
	statusMsg = "Get successful"
}
