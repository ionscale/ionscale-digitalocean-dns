package main

import (
	plugin "github.com/jsiebens/libdns-plugin"
	"github.com/libdns/digitalocean"
)

func main() {
	plugin.Serve(&digitalocean.Provider{})
}
