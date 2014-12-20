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
	"crypto/rand"
	"crypto/sha512"
	"golang.org/x/crypto/curve25519"
	"net"
)

func createPrivateKey() [32]byte {
	var key [32]byte
	rand.Read(key[:])
	return key
}

func createPublicKey(privateKey [32]byte) (publicKey [32]byte) {
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	return publicKey
}

func hashPublicKey(publicKey []byte) []byte {
	x := sha512.Sum512(publicKey[:])
	y := sha512.Sum512(x[:])
	return y[0:16]
}

func publicKeyToIPv6(publicKey []byte) net.IP {
	hashedKey := hashPublicKey(publicKey[:])
	return net.IP.To16(hashedKey)
}

func isValidIPv6(ip []byte) bool {

	if ip == nil {
		return false
	}

	if ip[0] == 0xFC {
		return true
	}
	return false
}

type CryptoKeys struct {
	PublicKey, PrivateKey [32]byte
	IPv6                  net.IP
}

func generateKeys() CryptoKeys {

	keys := CryptoKeys{}

	for isValidIPv6(keys.IPv6) != true {
		keys.PrivateKey = createPrivateKey()
		keys.PublicKey = createPublicKey(keys.PrivateKey)
		keys.IPv6 = publicKeyToIPv6(keys.PublicKey[:])
	}

	return keys
}
