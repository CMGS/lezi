package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net"
	"path"
	"time"
)

func Publishv2(input chan []*FileEvent, registrar chan []*FileEvent, config *NetworkConfig) {
	rand.Seed(time.Now().UnixNano())
	for events := range input {
		for _, event := range events {
			addr := config.Servers[rand.Int()%len(config.Servers)]
			_, (*event.Fields)["name"] = path.Split(*event.Source)
			if err := udpStreamer(event, addr); err != nil {
				log.Println("Send event failed")
				continue
			}
		}
	}
}

func udpStreamer(logline *FileEvent, addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	writeJSON(conn, logline)
	return nil
}

func writeJSON(w io.Writer, logline *FileEvent) {
	encoder := json.NewEncoder(w)
	encoder.Encode(logline)
}
