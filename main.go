package main

import (
	"campaign-service/internal/db"
	"campaign-service/internal/server"
	"campaign-service/util"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// @title Service API
// @version 1.0

// @host localhost:8089
// @BasePath /
// @schemes http

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	cfg, err := util.LoadConfig(".", *environment)
	if err != nil {
		println("load config failed: ", err)
	}

	dbIns, err := db.Connect(cfg)
	if err != nil {
		println("connect to database failed: ", err)

	}

	serverIns, err := server.NewServer(cfg, dbIns)
	if err != nil {
		println("failed to create server: ", err)
	}

	go func() {
		err = serverIns.Run()
		if err != nil {
			println("failed to create server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit

	println("Received signal", logrus.Fields{"signal": sig})
	serverIns.Shutdown(sig)
}
