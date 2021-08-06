package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vinodborole/restapi-boilerplate/src/app/infra/logging"
)

//Logger is handler responsible for logging the REST request
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		if !strings.Contains(r.RequestURI, "/v1/runaholics/health") {
			log.Debugf(
				"%s %s %s %s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		}
	})
}

//LogInit initializes log for API request
func LogInit(r *http.Request, command string) (*logging.AuditLog, context.Context) {
	auditLog := logging.AuditLog{Request: &logging.Request{Command: command}}
	ctx := auditLog.LogMessageInit(r.Context())
	return &auditLog, ctx
}
