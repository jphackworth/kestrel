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
	"time"
)

func testMain() {
	dev, ifname, err := ip.OpenTun("")
	check(err)
	fmt.Printf("Device: %s, Interface: %s\n", dev, ifname)
	tun, err := netlink.LinkByName(ifname)
	check(err)
	err = netlink.LinkSetUp(tun)
	check(err)
	fmt.Printf("%s\n", tun)
	addr, err := netlink.ParseAddr("fc27:6ba1:f09b:6938:b0c0:8995:14ae:cfb1/8")
	check(err)
	netlink.AddrAdd(tun, addr)
	netlink.LinkSetMTU(tun, 1312)
	err = netlink.LinkSetUp(tun)
	check(err)
	time.Sleep(60000 * time.Millisecond)
}
