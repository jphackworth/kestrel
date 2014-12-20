zlarkd
======

An experimental cjdns router implementation.

## Background

Cjdns is the name of a [meshnet routing protocol](https://github.com/cjdelisle/cjdns/blob/master/doc/Whitepaper.md) and [router implementation](https://github.com/cjdelisle/cjdns) from [cjd](https://github.com/cjdelisle).

As of November 2014, the main developer (cjd) stepped down without leaving a successor. For users of the software, this leaves the following options:

1. Stop using cjdns
2. Use cjdns until it no longer works
3. Maintain the reference cjdns router implementation
4. Develop a new cjdns router

Zlarkd is my attempt at option #4. I believe if we build simpler cjdns router implementations, we will be better able to debug, enhance and grow cjdns networks (like Hyperboria) in the future.

## Known Issues/Limitations

1. I'm writing Zlarkd in Golang. I have no experience with Golang prior to this project
2. Cjdns is an encrypted network protocol. I haven't implemented encrypted network protocols  
2. Zlarkd is being written for Linux. Little effort will be made to make Zlarkd portable to non-x86 architectures, or operating systems other than Linux. If you need multiplatform support, stick with cjdroute. If you would like to help make zlarkd work on other platforms, please get in contact.

## FAQ

**Q: Is this a fork of cjdns?**

*A: No. This is my attempt to write an interoperable cjdns router in golang.*

**Q: What is the status of Zlarkd?** 

*A: Zlarkd is not ready for testing or production use*

**Q: Why not work on cjdroute instead?**

*A: I felt like learning Golang rather than adopting someone else's codebase*