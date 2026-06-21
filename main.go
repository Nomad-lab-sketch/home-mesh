package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

const serviceName = "_home-mesh._tcp"

func main() {
	host, _ := os.Hostname()
	name := fmt.Sprintf("%s-%d", host, os.Getpid())

	info := []string{"this is golang agent"}

	service, err := mdns.NewMDNSService(
		name,
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

	fmt.Println("Published as:", name)

	entriesCh := make(chan *mdns.ServiceEntry, 16)

	entryName := name + "." + serviceName + ".local."

	go func() {
		for entry := range entriesCh {

			if entry.Name == entryName {
				continue // игнорируем себя
			}

			fmt.Printf(
				"Found: %s %v %d\n",
				entry.Name,
				entry.AddrV4,
				entry.Port,
			)
		}
	}()

	go func() {
		for {
			mdns.Lookup(serviceName, entriesCh)
			time.Sleep(10 * time.Second)
		}
	}()

	select {}
}
