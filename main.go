package main

import (
	"dnsproxy/pkg/localip"
	"log"

	"github.com/miekg/dns"
)

func main() {

	localip.Getip()

	// Create a DNS handler function
	dnsHandler := func(w dns.ResponseWriter, r *dns.Msg) {
		// Create a UDP client to communicate with the upstream DNS server
		client := &dns.Client{Net: "udp"}

		// Configure the upstream DNS server address and port
		upstreamAddr := "8.8.8.8:53"

		// Send the DNS query to the upstream DNS server
		response, _, err := client.Exchange(r, upstreamAddr)
		if err != nil {
			log.Printf("Failed to send DNS query: %s", err.Error())
			return
		}

		// Send the DNS response back to the client
		err = w.WriteMsg(response)
		if err != nil {
			log.Printf("Failed to send DNS response: %s", err.Error())
		}
	}

	// Start the DNS server
	server := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: dns.HandlerFunc(dnsHandler),
	}

	log.Printf("DNS proxy server listening on :53...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
