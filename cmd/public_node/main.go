package main

import (
	"context"
	"github.com/telf01/soo/pkg/configurator"
	"github.com/telf01/soo/pkg/logger"
	"github.com/telf01/soo/pkg/public_node/auth"
	"github.com/telf01/soo/pkg/public_node/network"
	"github.com/telf01/soo/pkg/public_node/node"
	"github.com/telf01/soo/pkg/public_node/persistence/auth_db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Save data from config.json to configurator.Cfg object.
	err := configurator.LoadSettings("config.json")
	if err != nil {
		panic("Can't initialize logger: " + err.Error())
	}

	// Create logger.L object
	logger.Initialize()

	// Initialize auth database connection.
	DB, err := auth_db.NewDB("auth_db.db")
	if err != nil {
		logger.L.Sugar().Fatal("Can't connect to database.")
	}
	a := auth.NewAuth(DB)

	// Initialize network
	s := http.Server{}
	net := network.NewNetwork(configurator.Cfg.NetworkConfiguration.Address, &s)
	cm := node.NewConnectionManager(net, a)
	cmp := node.NewComposer(cm)
	cmp.Start()

	// Graceful shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	log.Println("Got signal:", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
