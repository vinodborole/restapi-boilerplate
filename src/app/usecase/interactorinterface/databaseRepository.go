package interactorinterface

import (
	"app/infra/database"
	"github.com/jinzhu/gorm"
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
