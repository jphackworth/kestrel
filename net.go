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
	"github.com/vishvananda/netlink"
	"log"
	"net"
	"os"
)

type Server struct {
	Conn    *net.UDPConn
	TunFile *os.File
	TunLink netlink.Link
	//observers map[string][]data.EventsChannel
}

func startUDPServer(listenAddress string, server *Server) *Server {
	addr, err := net.ResolveUDPAddr("udp4", listenAddress)
	check(err)

	conn, err := net.ListenUDP("udp4", addr)
	check(err)

	server.Conn = conn
	//observers: make(map[string][]data.EventsChannel),

	log.Printf("Listening on %s\n", listenAddress)
	go server.ReceiveDatagrams()
	return server
}

func (self *Server) ReceiveDatagrams() {

	for {
		buffer := make([]byte, 1024)

		if c, addr, err := self.Conn.ReadFromUDP(buffer); err != nil {

			log.Printf("Oops! %d byte datagram from %s with error %s\n", c, addr.String(), err.Error())
			return

		} else {
			log.Printf("%d byte datagram from %s\n", c, addr.String())
		}
	}
	panic("should never have got myself into this.")
}
