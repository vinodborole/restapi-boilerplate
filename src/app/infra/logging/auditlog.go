package logging

import (
	"app/gateway/appcontext"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

//AuditLog Constants indicating the commands status on the Rest Server
const (
	RCVD      = "Received"
	COMPLETED = "Completed"
	FAILED    = "Failed"
)

//Request captures the Request Command and its parameters
type Request struct {
	Command string                 `json:"cmd"`
	Params  map[string]interface{} `json:"params"`
}

//AuditLog is the receiver object providing the Logging functionality
type AuditLog struct {
	Request   *Request
	StartTime time.Time
	Logger    *logrus.Entry
	ReqID     string
	UserID    uint
	ctx       context.Context
	User      string
}

//LogMessageInit initializes the AuditLog and setups the logger with the Request
func (alog *AuditLog) LogMessageInit(ctx context.Context) context.Context {
	alog.ReqID = uuid.New().String()
	Logger, ctx := appcontext.LoggerAndContext(ctx, alog.ReqID)
	alog.ctx = ctx
	alog.Logger = Logger.WithFields(logrus.Fields{
		"request": alog.Request,
	})
	return ctx
}

//LogMessageReceived to be invoked when the Request messages is recieved on the server.
//It will logs the message to log file and audit log
func (alog *AuditLog) LogMessageReceived() {
	alog.Logger.Infoln(RCVD)
	alog.logStartAudit(RCVD)
	//RequestBytes, _ := json.Marshal(alog.Request.Params)
}

func (alog *AuditLog) logStartAudit(status string) {
	alog.StartTime = time.Now()
	//RequestBytes, _ := json.Marshal(alog.Request.Params)
}
func (alog *AuditLog) logEndAudit(status string) {
	//duration := time.Since(alog.StartTime)
}

//LogMessageEnd to indicate the completion of command.
func (alog *AuditLog) LogMessageEnd(success *bool, statusMsg *string) {
	if *statusMsg != "" {
		alog.Logger = alog.Logger.WithField("reason", statusMsg)
	}
	if *success {
		alog.Logger.Infoln(COMPLETED)
		alog.logEndAudit(COMPLETED)
	} else {
		alog.Logger.Infoln(FAILED)
		alog.Logger.Errorln(FAILED)
		alog.logEndAudit(FAILED)
	}
}
