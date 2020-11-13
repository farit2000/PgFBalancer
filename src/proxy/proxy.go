package proxy

import (
	"gopkg.in/gcfg.v1"
	"log"
	"net"
)

type config struct {
	PgReplicaProxy struct {
		Listen  []string
		Backend []string
	}
}

var masterRequestChannel = make(chan serverRequest)
var replicaRequestChannel = make(chan serverRequest)
var serverStatusUpdateChannel = make(chan serverStatusUpdate)
var exitChan = make(chan bool)

func StartProxy() {
	cfg := config{}
	err := gcfg.ReadFileInto(&cfg, "PgFBalancer.cfg")
	if err != nil {
		log.Fatal(err)
	}

	go serverStatusOracle()
	go manageBackendKeyDataStorage()
	for _, backend := range cfg.PgReplicaProxy.Backend {
		go monitorBackend(backend)
	}
	for _, listen := range cfg.PgReplicaProxy.Listen {
		go listenFrontend(listen)
	}

	// Don't finish Proxy()
	<-exitChan
}

func listenFrontend(listen string) {
	ln, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleIncomingConnection(conn, masterRequestChannel, replicaRequestChannel)
	}
}
