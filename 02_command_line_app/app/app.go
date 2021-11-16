package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate will return the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command-line Application"
	app.Usage = "Get ip address from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IP of host",
			Flags:  flags,
			Action: findIpAddr,
		},
		{
			Name:   "ns",
			Usage:  "Find the NS of host",
			Flags:  flags,
			Action: findNameServers,
		},
	}

	return app
}

func findIpAddr(c *cli.Context) {
	host := c.String("host")

	fmt.Println("Finding IPs for host: " + host)
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findNameServers(c *cli.Context) {
	host := c.String("host")

	fmt.Println("Finding NS for host: " + host)
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
