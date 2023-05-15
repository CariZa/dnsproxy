package localip

import (
	"fmt"
	"net"
)

func Getip() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Failed to retrieve network interfaces: %s", err.Error())
		return
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Failed to retrieve addresses for interface %s: %s\n", iface.Name, err.Error())
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			if ip.IsLoopback() {
				continue
			}

			if ip.To4() != nil {
				fmt.Printf("IPv4: %s\n", ip.String())
			}
		}
	}
}
