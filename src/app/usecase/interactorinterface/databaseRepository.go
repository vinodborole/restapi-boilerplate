package interactorinterface

import (
	"github.com/jinzhu/gorm"
	"github.com/vinodborole/restapi-boilerplate/src/app/infra/database"
)

//DatabaseRepository database repository
type DatabaseRepository interface {
	//Transaction Operations
	OpenTransaction() (*gorm.DB, error)
	CommitTransaction(transaction *gorm.DB) error
	RollBackTransaction(transaction *gorm.DB) error

	//App Info
	GetApp(appName string) (database.App, error)
	CreateApp(app *database.App) error
}
