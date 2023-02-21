package app

import (
	"github.com/urfave/cli"
)

func StartApp() *cli.App {
	app := cli.NewApp()
	app.Name = "s3arch of ips v3ronez"
	app.Usage = "Search for ip of a domain name"

	app.Commands = []cli.Command{
		{
			Name:  "search-ip",
			Usage: "search an ip from a domain",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "amazon.com.br",
				},
			},
			Action: searchIp,
		},
	}
	return app
}
