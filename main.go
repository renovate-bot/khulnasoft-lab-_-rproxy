package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/khulnasoft-lab/rproxy/auth"
	"github.com/khulnasoft-lab/rproxy/config"
	"github.com/khulnasoft-lab/rproxy/proxy"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func start(c *cli.Context) error {
	port := c.Int("port")
	upstream := c.String("upstream")
	configPath := c.String("config")

	config, err := config.Parse(&configPath)
	if err != nil {
		log.Fatalf("Could not read configuration file: %v", err)
		return err
	}

	upstreamURL, err := url.Parse(upstream)
	if err != nil {
		log.Fatalf("Could not parse upstream URL: %v", err)
		return err
	}
	reverseProxy := proxy.NewReverseProxy(upstreamURL)
	http.HandleFunc("/", auth.Middleware(reverseProxy.ServeHTTP, *config))

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Proxy could not start %v", err)
		return err
	}

	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Proxy could not start %v", err)
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Rproxy"
	app.Usage = "Reverse Proxy with Basic Authentication"
	app.Version = version
	app.Authors = []*cli.Author{
		{
			Name: "KhulnaSoft Labs.",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "Start the proxy",
			Action: start,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "port",
					Usage: "Port used by the proxy",
					Value: 6634,
				}, &cli.StringFlag{
					Name:  "upstream",
					Usage: "Upstream server to proxy to",
					Value: "https://httpbin.org",
				}, &cli.StringFlag{
					Name:  "config",
					Usage: "Configuration file path",
					Value: "config.yaml",
				},
			},
		},
	}
	app.Run(os.Args)
}
