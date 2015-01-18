// Copyright 2014 JPH <jph@hackworth.be>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/adampresley/sigint"
	"log"
	"os"
	"time"
)

// http://guzalexander.com/2013/12/06/golang-channels-tutorial.html

// serverInfo struct contains:
// - tun device file handle,
// - udp network server connection
// - logger handle
// - goroutine descriptors
//
// These allow kestrel to shutdown() gracefully when it receives an
// interrupt or fatal erro
var server *ServerInfo

// config struct contains the variables from the main kestrel
// configuration file
//var config tomlConfig

//func start(config tomlConfig) {
func run(configFile string) {
	sleepInterval := 60

	config := readConfigFile(configFile)
	log.Printf("Starting\n")

	//startUDPServer(config.Server.Listen, &server)
	//initTunDevice(config, &server)

	//server = startServer(config)
	router := newRouter(&config)
	router.Start()

	for {
		log.Printf("Sleeping for %d seconds\n", sleepInterval)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}

// func shutdown() {
// 	log.Printf("Shutting down\n")
// 	// log.Printf("Closing UDP connection... ")
// 	// //server.Tun.file
// 	// err := server.Conn.Close()
// 	// if err != nil {
// 	// 	log.Printf("Error closing UDP connection: %s\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// log.Printf("Closing %s", server.Tun.Name)
// 	// if err != nil {
// 	// 	log.Printf("Error closing %s: %s", server.Tun.Name, err)
// 	// }
// 	log.Printf("See you next time!\n")
// 	os.Exit(0)
// }

//ctrl-c interrupt code from http://adampresley.com/2014/12/15/handling-ctrl-c-in-go-command-line-applications.html
func main() {

	sigint.ListenForSIGINT(func() {
		log.Printf("Received SIGINT.\n")
		server.Shutdown()
	})

	app := getApp()
	app.Run(os.Args)

}
