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
	"github.com/coreos/flannel/pkg/ip"
	"github.com/vishvananda/netlink"
	"log"
	"os"
)

func initTunDevice(config tomlConfig, server *Server) {

	server.TunFile, server.TunLink = openDevice(config.Server.Device)
	configureDevice(server.TunLink, config.Server.IPv6, 1312)

	return
}

// openDevice opens the specified device and sets it to UP status.
//
// The file handle to the opened tun device is returned so that it
// can be closed in the future, and the tun device as netlink.Link so
// we can configure it (MTU, IPv6 address) via the netlink package.
//
// TODO:
// - Add closeDevice function
// - Create our own ip.OpenTun since we dont use much else from github.com/coreos/flannel/pkg/ip
func openDevice(deviceName string) (*os.File, netlink.Link) {
	log.Printf("Opening %s", deviceName)
	dev, ifname, err := ip.OpenTun(deviceName)
	check(err)
	tun, err := netlink.LinkByName(ifname)
	check(err)
	log.Printf("Bringing %s UP", ifname)
	err = netlink.LinkSetUp(tun)
	check(err)
	return dev, tun
}

// Adds the IPv6 address and sets the MTU for specified tun device
func configureDevice(tun netlink.Link, ipv6 string, mtu int) {
	addDeviceAddress(tun, ipv6)
	setDeviceMTU(tun, mtu)
}

// Add the ipv6 address on specified tun device
//
// TODO: test behavior when the same address is added twice
func addDeviceAddress(tun netlink.Link, ipv6 string) {
	addr, err := netlink.ParseAddr(fmt.Sprintf("%s/8", ipv6))
	check(err)
	log.Printf("Assigning %s to %s", ipv6, tun.Attrs().Name)
	err = netlink.AddrAdd(tun, addr)
	check(err)
}

func setDeviceMTU(tun netlink.Link, mtu int) {

	log.Printf("Setting %s MTU to %d", tun.Attrs().Name, mtu)
	err := netlink.LinkSetMTU(tun, mtu)
	check(err)
}

// func testMain() {
// 	dev, ifname, err := ip.OpenTun("")
// 	check(err)
// 	fmt.Printf("Device: %s, Interface: %s\n", dev, ifname)
// 	tun, err := netlink.LinkByName(ifname)
// 	check(err)
// 	err = netlink.LinkSetUp(tun)
// 	check(err)
// 	fmt.Printf("%s\n", tun)
// 	addr, err := netlink.ParseAddr("fc27:6ba1:f09b:6938:b0c0:8995:14ae:cfb1/8")
// 	check(err)
// 	netlink.AddrAdd(tun, addr)
// 	netlink.LinkSetMTU(tun, 1312)
// 	err = netlink.LinkSetUp(tun)
// 	check(err)
// 	time.Sleep(60000 * time.Millisecond)
// }
