package main

import (
	"log"
	"os"
)

// func startServer(c TomlConfig) *ServerInfo {
// 	s := &ServerInfo{Server: startUDPServer(c.Server.Listen),
// 		TunDevice: startTunDevice(c)}

// 	return s
// }

func (s *ServerInfo) Shutdown() {
	log.Printf("shutting down")
	os.Exit(0)
}
