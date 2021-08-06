package usecase

import (
	"context"
	"sync"

	InteractorIntf "github.com/vinodborole/restapi-boilerplate/src/app/usecase/interactorinterface"
)

//Interactor is the receiver object
type Interactor struct {
	Db         InteractorIntf.DatabaseRepository
	AppContext context.Context
	DBMutex    sync.Mutex
	Refresh    bool
}
