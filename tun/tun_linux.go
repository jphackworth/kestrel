package tun

import (
	_ "fmt"
	"github.com/vishvananda/netlink"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

const (
	IFF_NO_PI = 0x10
	IFF_TUN   = 0x01
	IFF_TAP   = 0x02
	TUNSETIFF = 0x400454CA
)

func (tun *Tun) Open() {
	deviceFile := "/dev/net/tun"
	fd, err := os.OpenFile(deviceFile, os.O_RDWR, 0)
	if err != nil {
		log.Fatalf("[CRIT] Note: Cannot open TUN/TAP dev %s: %v", deviceFile, err)
	}
	tun.fd = fd
	//tun.link = netlink.LinkByName(string(ifr))

	ifr := make([]byte, 18)
	ifr[17] = IFF_NO_PI
	ifr[16] = IFF_TUN
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(tun.fd.Fd()), uintptr(TUNSETIFF),
		uintptr(unsafe.Pointer(&ifr[0])))
	if errno != 0 {
		log.Fatalf("[CRIT] Cannot ioctl TUNSETIFF: %v", errno)
	}

	tun.actualName = string(ifr)
	tun.actualName = tun.actualName[:strings.Index(tun.actualName, "\000")]
	tun.link, err = netlink.LinkByName(tun.actualName)
	check(err)
	log.Printf("[INFO] TUN/TAP device %s opened.", tun.actualName)
}

func (tun *Tun) SetupAddress(newAddr string, newMTU int) {

	cmd1 := exec.Command("/sbin/ip", "link", "set", "dev", tun.actualName, "up")
	log.Printf("[DEBG] ip(8) command: %v", strings.Join(cmd1.Args, " "))
	err := cmd1.Run()
	check(err)

	log.Printf("mtu is %s", strconv.Itoa(newMTU))

	cmd2 := exec.Command("/sbin/ip", "link", "set", "dev", tun.actualName, "mtu", strconv.Itoa(newMTU))
	log.Printf("[DEBG] ip(8) command: %v", strings.Join(cmd2.Args, " "))
	err = cmd2.Run()
	check(err)

	cmd3 := exec.Command("/sbin/ip", "-6", "addr", "add", newAddr, "dev", tun.actualName)
	log.Printf("[DEBG] ip(8) command: %v", strings.Join(cmd3.Args, " "))
	err = cmd3.Run()
	check(err)

	//tun.AddAddress(newAddr)
	//tun.SetMTU(newMTU)
}

//func (tun *Tun) AddAddress(newAddr string) {
// 	addr_with_mask := fmt.Sprintf("%s/8", newAddr)
// 	fmt.Println(addr_with_mask)

// 	addr, err := netlink.ParseAddr(fmt.Sprintf("%s/8", newAddr))
// 	if err != nil {
// 		log.Fatalf("Cannot parse address: %s", addr_with_mask)
// 	}
// 	fmt.Printf("Assigning address to %s\n", tun.link.Attrs().Name)
// 	log.Printf("Assigning %s to %s", addr_with_mask, tun.link.Attrs().Name)
// 	err = netlink.AddrAdd(tun.link, addr)
// 	check(err)
// }

// func (tun *Tun) SetMTU(newMTU int) {
// 	log.Printf("Setting %s MTU to %d", tun.link.Attrs().Name, newMTU)
// 	err := netlink.LinkSetMTU(tun.link, newMTU)
// 	check(err)
// }

// func (tun *Tun) SetupAddress2(addr, mask string) {
// 	cmd := exec.Command("ifconfig", tun.actualName, addr,
// 		"netmask", mask, "mtu", "1500")
// 	log.Printf("[DEBG] ifconfig command: %v", strings.Join(cmd.Args, " "))
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Printf("[EROR] Linux ifconfig failed: %v.", err)
// 	}
// }
