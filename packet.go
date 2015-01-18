// packet.go

package main

import (
	"code.google.com/p/gopacket"
	_ "code.google.com/p/gopacket/layers"
	"code.google.com/p/gopacket/pcap"
	_ "log"
)

// Used to create a packet capture source for tun device
func newPacketSource(deviceName, filter string) *gopacket.PacketSource {

	if handle, err := pcap.OpenLive(deviceName, 1600, true, 0); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter(filter); err != nil { // optional
		panic(err)
	} else {
		return gopacket.NewPacketSource(handle, handle.LinkType())
	}

}
