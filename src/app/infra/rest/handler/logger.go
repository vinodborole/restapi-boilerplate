package handler

import (
	"app/infra/logging"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
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
