package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

const serviceName = "_homemesh._tcp"

func main() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{"agent"}
	service, err := mdns.NewMDNSService(
		host,
		serviceName,
		"",
		"",
		8000,
		nil,
		info,
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Fatal(err)
	}
	defer server.Shutdown()

	fmt.Println("mDNS service published")

	entriesCh := make(chan *mdns.ServiceEntry, 10)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("%s (%s:%d)\n", entry.Name, entry.AddrV4[0], entry.Port)
			// TODO handle entry
		}
	}()

	for {
		mdns.Lookup(serviceName, entriesCh)
		time.Sleep(10 * time.Second)
	}
}
