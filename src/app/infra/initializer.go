package infra

import (
	"sync"

	"github.com/vinodborole/restapi-boilerplate/src/app/gateway"
	"github.com/vinodborole/restapi-boilerplate/src/app/infra/database"
	"github.com/vinodborole/restapi-boilerplate/src/app/usecase"
)

//UsecaseInteractor  provides receiver object for services
var UsecaseInteractor *usecase.Interactor
var once sync.Once

//GetUseCaseInteractor gets a singleton instance of DeviceInteractor with DatabaseRepository
func GetUseCaseInteractor() *usecase.Interactor {
	once.Do(func() {
		DatabaseRepository := gateway.DatabaseRepository{Database: database.GetWorkingInstance()}
		UsecaseInteractor = &usecase.Interactor{Db: &DatabaseRepository}
	})
	return UsecaseInteractor
}
