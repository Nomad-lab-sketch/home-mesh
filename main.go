package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/mdns"
)

func main() {

	host, error := os.Hostname()

	if error == nil {
		fmt.Printf("Hostname: %s", host)
		return
	}

	info := []string{"My awesome service"}

	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)

	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}
