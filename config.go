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
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"net"
)

type tomlConfig struct {
	Server serverInfo
}

type serverInfo struct {
	Listen     string `toml:"listen"`
	PublicKey  string `toml:"public_key"`
	PrivateKey string `toml:"private_key"`
	IPv6       string `toml:"ipv6"`
}

func getApp() *cli.App {
	app := cli.NewApp()
	app.Name = "zlarkd"
	app.Usage = "An experimental cjdns router"
	app.Author = "JPH"
	app.Email = "jph@hackworth.be"
	app.Version = "0.1.0a"
	app.Commands = []cli.Command{
		{
			Name:  "genconf",
			Usage: "generate a new config",
			Action: func(c *cli.Context) {
				generateConfig()
			},
		}, {
			Name:  "readconf",
			Usage: "read a specified config",
			Action: func(c *cli.Context) {
				readConfig(c.String("config"))
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "/etc/zlarkd/zlarkd.toml",
			Usage: "Config file location",
		},
	}

	return app
}

func readConfig(configPath string) {

}

func generateConfig() {
	conf := tomlConfig{}
	keys := generateKeys()
	conf.Server.IPv6 = keys.IPv6.String()
	conf.Server.PublicKey = fmt.Sprintf("%s.k", base32Encode(keys.PublicKey[:])[:52])
	conf.Server.PrivateKey = hex.EncodeToString(keys.PrivateKey[:])
	conf.Server.Listen = generateListenAddress()
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(conf)
	check(err)
	fmt.Println(buf.String())
}

func generateListenAddress() (listenAddress string) {
	addr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:0")
	check(err)
	conn, err := net.ListenUDP("udp4", addr)
	check(err)
	listenAddress = conn.LocalAddr().String()
	conn.Close()
	return listenAddress

}
