package converter

import (
	"github.com/vinodborole/restapi-boilerplate/src/app/infra/database"
	Restmodel "github.com/vinodborole/restapi-boilerplate/src/app/infra/rest/generated/server/go"
)

//GetFitnessAppSwaggerResponse get fitness app swagger response
func GetAppSwaggerResponse(dbFitnessApp *database.App) (Restmodel.App, error) {
	var App Restmodel.App
	App.Description = dbFitnessApp.Description
	App.Name = dbFitnessApp.Name
	App.Port = dbFitnessApp.Port
	App.Url = dbFitnessApp.Url
	return App, nil
}
