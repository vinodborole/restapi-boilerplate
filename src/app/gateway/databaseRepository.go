package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/vinodborole/restapi-boilerplate/src/app/infra/database"
)

//DatabaseRepository represents the Application Database Repository
type DatabaseRepository struct {
	Database *database.Database
}

//GetDBHandle returns a handle to DB, using which the database operations can be performed
func (dbRepo *DatabaseRepository) GetDBHandle() *gorm.DB {
	return dbRepo.Database.Instance
}

//OpenTransaction begins the database transaction
func (dbRepo *DatabaseRepository) OpenTransaction() (*gorm.DB, error) {
	transaction := dbRepo.GetDBHandle().Begin()
	return transaction, transaction.Error
}

//CommitTransaction commits the database transaction
func (dbRepo *DatabaseRepository) CommitTransaction(transaction *gorm.DB) error {
	err := transaction.Commit().Error
	if err != nil {
		panic(err)
	}
	transaction = nil
	return err
}

//RollBackTransaction rolls back the database transaction
func (dbRepo *DatabaseRepository) RollBackTransaction(transaction *gorm.DB) error {
	err := transaction.Rollback().Error
	transaction = nil
	return err
}
