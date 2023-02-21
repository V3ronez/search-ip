package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func searchIp(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
