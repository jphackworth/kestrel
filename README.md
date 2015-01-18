Kestrel
=======

An experimental cjdns router implementation.

## Background

Cjdns is the name of a [meshnet routing protocol](https://github.com/cjdelisle/cjdns/blob/master/doc/Whitepaper.md) and [router implementation](https://github.com/cjdelisle/cjdns) from [cjd](https://github.com/cjdelisle).

As of November 2014, the main developer (cjd) stepped down without leaving a successor. For users of the software, this leaves the following options:

1. Stop using cjdns
2. Use cjdns until it no longer works
3. Maintain the reference cjdns router implementation
4. Develop a new cjdns router

kestrel is my attempt at option #4. I believe if we build simpler cjdns router implementations, we will be better able to debug, enhance and grow cjdns networks (like Hyperboria) in the future.

## Known Issues/Limitations

1. I'm writing Kestrel in Golang. I have no experience with Golang prior to this project
2. Cjdns is an encrypted network protocol. I haven't implemented encrypted network protocols  
2. kestrel is being written for Linux. Little effort will be made to make kestrel portable to non-x86 architectures, or operating systems other than Linux. If you need multiplatform support, stick with cjdroute. If you would like to help make kestrel work on other platforms, please get in contact.

## FAQ

**Q: Is this a fork of cjdns?**

*A: No. This is my attempt to write an interoperable cjdns router in golang.*

**Q: What is the status of kestrel?** 

*A: kestrel is not ready for testing or production use*

**Q: Why not work on cjdroute instead?**

*A: I felt like learning Golang rather than adopting someone else's codebase*

## For Developers

If you want to test/hack on it in its current state:

1. Install/setup Go: https://golang.org/doc/install
2. go get github.com/jphackworth/kestrel
3. go build github.com/jphackworth/kestrel
4. go install github.com/jphackworth/kestrel
5. $GOPATH/bin/kestrel

NOTE: There's not much to see at the moment.
