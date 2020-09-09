package usecase

import (
	InteractorIntf "app/usecase/interactorinterface"
	"context"
	"sync"
)

//Interactor is the receiver object
type Interactor struct {
	Db         InteractorIntf.DatabaseRepository
	AppContext context.Context
	DBMutex    sync.Mutex
	Refresh    bool
}
