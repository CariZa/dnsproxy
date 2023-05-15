# DNSProxy POC

POC for a DNSProxy concept

Forwarding DNS requests from a DNS server not running on port 53, so that resolv.conf can use this DNSPRoxy that will run on the default port 53.

To run the POC

    $ go run main.go