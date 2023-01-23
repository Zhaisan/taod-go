package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	author "restapi/internal/author/db"
	config "restapi/internal/config"
	"restapi/internal/user"
	"restapi/pkg/client/postgresql"
	"restapi/pkg/logging"
	"time"
)




func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New() //создаем роутер

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	repository := author.NewRepository(postgreSQLClient, logger)

	one, err := repository.FindOne(context.TODO(), "bc97b486-8ce0-4515-8d1b-95327149dfa0")
	if err != nil {
		return
	}
	logger.Infof("%v", one)

	all, err := repository.FindAll(context.TODO())
	if err != nil {
		logger.Fatalf("%v", err)
	}

	for _, ath := range all {
		logger.Infof("%v", ath)
	}

	logger.Info("register user handler")
	handler := user.NewHandler(logger) //создаем хендлер
	handler.Register(router) //зарегали хендлер в роутере

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {

	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil{
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s: %s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}


	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	} //ссылочка на сервер


	logger.Fatal(server.Serve(listener))
}
