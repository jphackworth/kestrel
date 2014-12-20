//taken from https://gist.github.com/kylelemons/4585718

package main

import (
	"fmt"
)

var i2b = []byte("0123456789bcdfghjklmnpqrstuvwxyz")
var b2i = func() []byte {
	var ascii [256]byte
	for i := range ascii {
		ascii[i] = 255
	}
	for i, b := range i2b {
		ascii[b] = byte(i)
	}
	return ascii[:]
}()

func base32Encode(in []byte) (out []byte) {
	var wide, bits uint
	for len(in) > 0 {
		// Add the 8 bits of data from the next `in` byte above the existing bits
		wide, in, bits = wide|uint(in[0])<<bits, in[1:], bits+8
		for bits > 5 {
			// Remove the least significant 5 bits and add their character to out
			wide, out, bits = wide>>5, append(out, i2b[int(wide&0x1F)]), bits-5
		}
	}
	// If it wasn't a precise multiple of 40 bits, add some padding based on the remaining bits
	if bits > 0 {
		out = append(out, i2b[int(wide)])
		out = append(out, "====="[bits:]...)
	}
	return out
}

func base32Decode(in []byte) (out []byte, err error) {
	var wide, bits uint
	for len(in) > 0 && in[0] != '=' {
		// Add the 5 bits of data corresponding to the next `in` character above existing bits
		wide, in, bits = wide|uint(b2i[int(in[0])])<<bits, in[1:], bits+5
		if bits >= 8 {
			// Remove the least significant 8 bits of data and add it to out
			wide, out, bits = wide>>8, append(out, byte(wide)), bits-8
		}
	}
	// If there was padding, there will be bits left, but they should be zero
	if wide != 0 {
		return nil, fmt.Errorf("extra data at end of decode")
	}

	return out, nil
}
