package rest

import (
	"app/config"
	"app/infra"
	"app/infra/database"
	"app/infra/logging"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

//APIServer Receiver object for API Rest Server
type APIServer struct {
	http.Server
	shutdownReq chan bool
	wg          *sync.WaitGroup
	reqCount    uint32
}

//NewAPIServer provides an instance of APIServer
func NewAPIServer(wg *sync.WaitGroup) *APIServer {
	//create server
	s := &APIServer{
		Server: http.Server{
			Addr:         ":8080",
			ReadTimeout:  1000 * time.Second,
			WriteTimeout: 1000 * time.Second,
		},
		wg:          wg,
		shutdownReq: make(chan bool),
	}
	router := NewRouter()
	//set http server handler
	s.Handler = router
	router.HandleFunc("/shutdown", s.APIShutdownHandler)
	return s
}

//APIShutdownHandler provides shutdown handler for closing the REST Server
//It also closes the Database.
func (s *APIServer) APIShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))
	//First shut down the DB
	if err := database.GetWorkingInstance().Close(); err != nil {
		log.Errorln("Failed to shutdown the database engine", err)
	}
	//Do nothing if shutdown request already issued
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}
	go func() {
		s.shutdownReq <- true
	}()
}

//RunAPIServer starts the Database and Rest Server
func (s *APIServer) RunAPIServer() {
	//Start the server
	server := s
	//Initialize the DB
	user := config.GetConfig().MyAppConfig.Database.Username
	password := config.GetConfig().MyAppConfig.Database.Password
	dbName := config.GetConfig().MyAppConfig.Database.DBName
	location := config.GetConfig().MyAppConfig.Database.Location
	port := config.GetConfig().MyAppConfig.Database.Port
	log.Infoln("Setup Database")
	database.Setup(dbName, location, user, password, port)
	database.GetWorkingInstance().MigrateSchema()
	//Add Default Fabric from here
	alog := logging.AuditLog{Request: &logging.Request{Command: "default operations"}}
	ctx := alog.LogMessageInit(context.Background())
	alog.LogMessageReceived()
	success := true
	statusMsg := "default operations"
	setupDefaultData(ctx)
	log.Infoln("start server")
	alog.LogMessageEnd(&success, &statusMsg)
	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()
	//wait shutdown
	server.waitShutdown()
	<-done
	log.Printf("DONE!")
}
func (s *APIServer) waitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)
	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}
	log.Printf("Stoping http server ...")
	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
	s.wg.Done()
}

func setupDefaultData(ctx context.Context) {
	configObj := config.GetConfig()
	app := database.App{
		Name:        configObj.MyAppConfig.Application.Name,
		Description: configObj.MyAppConfig.Application.Description,
		Port:        configObj.MyAppConfig.Application.Port,
		Url:         "http://localhost",
	}
	_, err := infra.GetUseCaseInteractor().Db.GetApp(app.Name)
	if err != nil {
		log.Infoln("Creating app")
		err = infra.GetUseCaseInteractor().Db.CreateApp(&app)
		if err != nil {
			log.Errorln("Error creating App ", err.Error())
		}
	}
}
