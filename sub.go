package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gammazero/nexus/client"
	"github.com/gammazero/nexus/router"
	"github.com/gammazero/nexus/wamp"
	"./services/screenshot"
)
var logger = log.New(os.Stdout, "", 0)

const exampleTopic = "example.hello"

var cfg = client.Config{
  Realm:  "cheshire@anima-os.com",
  Logger: logger,
}

func register(nxr router.Router) {

	// Connect subscriber session.
	subscriber, err := client.ConnectLocal(nxr, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer subscriber.Close()

	// Define function to handle events received.
	evtHandler := func(args wamp.List, kwargs wamp.Dict, details wamp.Dict) {
		logger.Println("Received", exampleTopic, "event")
		if len(args) != 0 {
			logger.Println("  Event Message:", args[0])
		}
	}

	// Subscribe to topic.
	err = subscriber.Subscribe(exampleTopic, evtHandler, nil)
	if err != nil {
		logger.Fatal("subscribe error:", err)
	}
	logger.Println("Subscribed to", exampleTopic)

  // Register procedure "TakeScreenshot"
  procName := "TakeScreenshot"
  if err = subscriber.Register(procName, screenshot.TakeScreenshot, nil); err != nil {
    logger.Fatal("Failed to register procedure:", err)
  }
  logger.Println("Registered procedure", procName, "with router")

	// Wait for CTRL-c or client close while handling events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-subscriber.Done():
		logger.Print("Router gone, exiting")
		return // router gone, just exit
	}

	// Unsubscribe from topic.
	if err = subscriber.Unsubscribe(exampleTopic); err != nil {
		logger.Println("Failed to unsubscribe:", err)
	}
}
