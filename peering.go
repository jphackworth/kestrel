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
	"fmt"
	"net"
)

func UDPListen(listenAddress string) (localUDPServer *net.UDPConn) {
	localAddr, err := net.ResolveUDPAddr("udp", listenAddress)
	check(err)
	localUDPServer, err = net.ListenUDP("udp", localAddr)
	check(err)
	return localUDPServer
}

func UDPConnect(localServer *net.UDPConn, peerAddress string, key [32]byte, password string) {
	remoteAddr, err := net.ResolveUDPAddr("udp", peerAddress)
	check(err)
	localAddr, err := net.ResolveUDPAddr("udp", localServer.LocalAddr().String())
	check(err)
	con, err := net.DialUDP("udp", localAddr, remoteAddr)
	check(err)
	fmt.Println(con)
	con.Close()
}
