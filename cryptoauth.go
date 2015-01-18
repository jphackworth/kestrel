// cryptoauth.go

package main

import (
	"bytes"
	_ "code.google.com/p/gopacket"
	_ "code.google.com/p/gopacket/layers"
	"crypto/sha256"
	"encoding/binary"
	_ "encoding/hex"
	"fmt"
	"log"
)

const (
	CryptoHeader_MAXLEN = 120
)

type CryptoAuthHeader struct {
	Nonce     uint32
	Handshake CryptoAuth_Handshake
	Payload   []byte
}

// type CryptoAuthHeader struct {
// }

type CryptoAuth_Handshake struct {
	Stage         uint32
	Challenge     CryptoAuth_Challenge
	Nonce         [24]uint8 // 24 bytes
	PublicKey     [32]uint8
	Authenticator [16]uint8 // 16 bytes
	TempPublicKey [32]uint8 // 32 bytes
}

type CryptoAuth_Challenge struct {
	Type                                uint8
	Lookup                              [7]byte
	RequirePacketAuthAndDerivationCount uint16
	Additional                          uint16
}

func NewCryptoAuth(data []byte) *CryptoAuthHeader {
	var ca CryptoAuthHeader
	//var ca.Handshake
	r := bytes.NewReader(data)
	binary.Read(r, binary.BigEndian, &ca.Nonce)

	binary.Read(r, binary.BigEndian, &ca.Handshake.Stage)

	binary.Read(r, binary.BigEndian, &ca.Handshake.Challenge.Type)
	binary.Read(r, binary.BigEndian, &ca.Handshake.Challenge.Lookup)
	binary.Read(r, binary.BigEndian, &ca.Handshake.Challenge.RequirePacketAuthAndDerivationCount)
	binary.Read(r, binary.BigEndian, &ca.Handshake.Challenge.Additional)

	binary.Read(r, binary.BigEndian, &ca.Handshake.Nonce)
	binary.Read(r, binary.BigEndian, &ca.Handshake.PublicKey)
	binary.Read(r, binary.BigEndian, &ca.Handshake.Authenticator)
	binary.Read(r, binary.BigEndian, &ca.Handshake.TempPublicKey)

	binary.Read(r, binary.BigEndian, &ca.Payload)

	return &ca
}

func (c *CryptoAuthHeader) tryAuth() uint8 {

	//if
	return 0

}

// func (c *CryptoAuth_Challenge) getPasswordHash() (hash [32]uint8) {

// 	log.Printf("getPasswordHash: Challenge Type: %d\n", c.Type)

// 	if c.RequirePacketAuthAndDerivationCount {

// 		hash[0] ^= c.RequirePacketAuthAndDerivationCount[0]
// 		hash[0] ^= c.RequirePacketAuthAndDerivationCount[1]
// 	}
// }

func (h *CryptoAuth_Handshake) Dump() {
	fmt.Printf("Stage [%d], Nonce [%d], PublicKey [%s.k], Authenticator [%x], TempPublicKey [%x]\n",
		h.Stage, h.Nonce[:], base32Encode(h.PublicKey[:])[:52], h.Authenticator, h.TempPublicKey)
}

func (c *CryptoAuth_Challenge) Dump() {
	fmt.Printf("Type: %d", c.Type)
	fmt.Printf("Lookup: %d\n", c.Lookup[:])
	//fmt.Printf("RequirePacketAuthAndDerivationsCount: %s\n", hex.EncodeToString(c.RequirePacketAuthAndDerivationCount))
	fmt.Printf("Additional: %x\n", c.Additional)
	//	fmt.Println()
}

func (c *CryptoAuth_Challenge) hashPassword_256(password string) []byte {
	//var tempBuff [32]uint8
	//x := sha512.Sum512(publicKey[:])
	//h := sha256.New()
	pw_hash1 := sha256.Sum256([]byte(password))
	pw_hash2 := sha256.Sum256(pw_hash1[:32])
	//return h.Sum(h.Sum([]byte(password))[:32])[:12]
	return pw_hash2[:12]
}

func (c *CryptoAuth_Challenge) hashPassword(password string, authType int) {
	switch authType {
	case 1:
		c.hashPassword_256(password)
	default:
		log.Println("Error: hashPassword() Unsupported authType")
	}
}
