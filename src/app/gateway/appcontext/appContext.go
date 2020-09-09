package appcontext

import (
	"context"
	nlog "github.com/sirupsen/logrus"
	"runtime"
)

//ContextType represents type of the context, used to in the application
type ContextType int

const (
	//RequestIDKey represents the identifier of each REST request
	RequestIDKey ContextType = iota
	//UseCaseName usecase name
	UseCaseName
)

func getContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

//DecorateRuntimeContext represents a logger containing the "functionName" and "lineNumber" as fields
func DecorateRuntimeContext(logger *nlog.Entry) *nlog.Entry {
	if pc, _, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logger.WithField("line", line).WithField("func", fName)
	}
	return logger
}

//Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *nlog.Entry {
	newLogger := nlog.WithFields(nlog.Fields{
		"App": "myapp",
	})
	if ctx != nil {
		if ctxReqID, ok := ctx.Value(RequestIDKey).(string); ok {
			newLogger = newLogger.WithFields(nlog.Fields{
				"rqId": ctxReqID,
			})
		}
		if useCase, ok := ctx.Value(UseCaseName).(string); ok {
			newLogger = newLogger.WithFields(nlog.Fields{
				"UseCase": useCase,
			})
		}
	}
	return newLogger
}

//LoggerAndContext returns a zap logger and context
func LoggerAndContext(ctx context.Context, reqID string) (*nlog.Entry, context.Context) {
	reqCtx := getContext(ctx, reqID)
	logger := Logger(reqCtx)
	logger.Info("Request ID Created")
	return logger, reqCtx
}
