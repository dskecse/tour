package main

import (
    "fmt"
    "strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
    var ips []string
    for i := range ip {
        ips = append(ips, fmt.Sprintf("%v", ip[i]))
    }
    return fmt.Sprintf(strings.Join(ips, "."))
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
