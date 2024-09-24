package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type VMInfo struct {
	IPAddress  string `json:"ip-address"`
	MacAddress string `json:"mac-address"`
	Hostname   string `json:"hostname"`
	ClientID   string `json:"client-id"`
	ExpiryTime int64  `json:"expiry-time"`
}

func main() {
	statusFile := flag.String("f", "/var/lib/libvirt/dnsmasq/virbr0.status", "Path to the status file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: kvm-vm-iplookup [-f status_file] <vm_name>")
		os.Exit(1)
	}

	vmName := flag.Arg(0)

	data, err := os.ReadFile(*statusFile)
	if err != nil {
		fmt.Printf("Error reading status file: %v\n", err)
		os.Exit(1)
	}

	var vmInfos []VMInfo
	err = json.Unmarshal(data, &vmInfos)
	if err != nil {
		fmt.Printf("Error parsing JSON data: %v\n", err)
		os.Exit(1)
	}

	for _, info := range vmInfos {
		if info.Hostname == vmName {
			fmt.Println(info.IPAddress)
			os.Exit(0)
		}
	}

	fmt.Printf("VM '%s' not found in the status file\n", vmName)
	os.Exit(1)
}
