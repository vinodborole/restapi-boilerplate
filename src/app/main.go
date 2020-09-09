package main

import (
	"app/config"
	"app/infra/rest"
	"fmt"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello, my new app")
}
func init() {
	err := config.ReadYamlConfigFile()
	if err != nil {
		log.Errorln(err)
		return
	}
	logFile := config.GetConfig().MyAppConfig.Log.Location + config.GetConfig().MyAppConfig.Application.Name + ".log"
	LogWriter := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes
		MaxAge:     config.GetConfig().MyAppConfig.Log.MaxAge,
		MaxBackups: config.GetConfig().MyAppConfig.Log.MaxBackups,
		Compress:   true,
	}
	writerMap := lfshook.WriterMap{
		log.DebugLevel: &LogWriter,
		log.InfoLevel:  &LogWriter,
		log.WarnLevel:  &LogWriter,
		log.ErrorLevel: &LogWriter,
	}
	log.AddHook(lfshook.NewHook(
		writerMap,
		&log.JSONFormatter{},
	))
	log.SetOutput(ioutil.Discard)
	level := config.GetConfig().MyAppConfig.Log.Level
	if len(level) > 0 {
		switch level {
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warning":
			log.SetLevel(log.WarnLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		case "debug":
			log.SetLevel(log.DebugLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	} else {
		log.SetLevel(log.InfoLevel)
	}
	//startCron()
	startRestServer()
}
func startRestServer() error {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	//Wait Groups to start both the Rest Servers
	apiServer := rest.NewAPIServer(&wg)
	go apiServer.RunAPIServer()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Took ", elapsed)
	return nil
}
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
func startCron() {
	fmt.Println("starting cron..")
	doEvery(2*time.Hour, func(t time.Time) {
		fmt.Printf("%v: Cron Job to refresh fitness app tokens!\n", t)
	})
}
