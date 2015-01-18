package tun

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"log"
	"os"
)

type Tun struct {
	fd         *os.File
	link       netlink.Link
	actualName string
	ReadCh     chan []byte
	WriteCh    chan []byte
}

func (tun *Tun) Name() string {
	return tun.actualName
}

func New() *Tun {
	tun := &Tun{ReadCh: make(chan []byte), WriteCh: make(chan []byte)}
	return tun
}

func (tun *Tun) Start() {
	go tun.readLoop()
	go tun.writeLoop()
}

func (tun *Tun) writeLoop() {
	for {
		buf := <-tun.WriteCh
		_, err := tun.fd.Write(buf)
		if err != nil {
			log.Printf("[EROR] TUN/TAP: write failed: %v: %v", err, buf)
			tun.fd.Close()
			return
		}
	}
}

func (tun *Tun) ReadPacket() {
	var buf [4096]byte
	log.Println("about to read from fd")
	nread, err := tun.fd.Read(buf[:])
	//log.Println(nread)
	if nread > 0 {
		b := make([]byte, nread)
		copy(b, buf[:nread])
		fmt.Printf("x %x", b)
		tun.ReadCh <- b
	}
	if nread == 0 {
		log.Printf("closing tun fd")
		tun.fd.Close()
		return
	}
	if err != nil {
		log.Printf("[EROR] TUN/TAP: read failed: %v", err)
		tun.fd.Close()
		return
	}
}

func (tun *Tun) readLoop() {
	var buf [4096]byte
	for {
		nread, err := tun.fd.Read(buf[:])
		if nread > 0 {
			b := make([]byte, nread)
			copy(b, buf[:nread])
			fmt.Printf("x %x", b)
			tun.ReadCh <- b
		}
		if nread == 0 {
			log.Printf("closing tun fd")
			tun.fd.Close()
			return
		}
		if err != nil {
			log.Printf("[EROR] TUN/TAP: read failed: %v", err)
			tun.fd.Close()
			return
		}
	}
}

func (tun *Tun) Stop() {
	tun.fd.Close()
}
