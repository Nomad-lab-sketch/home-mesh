package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/mdns"
)

const serviceName = "_foobar._tcp"

func main() {
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("Hello"))
		})

		log.Println("HTTP server on :8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{"My awesome service"}
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

	select {}
}
