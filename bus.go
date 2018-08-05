package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gammazero/nexus/router"
	"github.com/gammazero/nexus/wamp"
    "net/http"
)

const address = "localhost:8080"

func check(r *http.Request) bool {
    return true
}


func main() {
	// Create router instance.
	routerConfig := &router.Config{
		RealmConfigs: []*router.RealmConfig{
			&router.RealmConfig{
				URI:           wamp.URI("cheshire@anima-os.com"),
				AnonymousAuth: true,
			},
		},
	}
	nxr, err := router.NewRouter(routerConfig, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer nxr.Close()

	// Create and run server
        s := router.NewWebsocketServer(nxr)
        s.Upgrader.CheckOrigin = check
	closer, err := s.ListenAndServe(address)
        
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Websocket server listening on ws://%s/", address)

	// Register internal subscribers and RPC services. 
	register(nxr)

	// Wait for SIGINT (CTRL-c), then close server and exit.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	closer.Close()
}
