package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // ensure hostname is passed
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <hostname>")
        os.Exit(1)
    }

    hostname := os.Args[1]

    // DNS lookup
    ips, err := net.LookupIP(hostname)
    if err != nil {
        fmt.Printf("Error looking up IP for %s: %v\n", hostname, err)
        os.Exit(1)
    }

    // Output the resolved IP addresses.
    fmt.Printf("IP addresses for %s:\n", hostname)
    for _, ip := range ips {
        fmt.Println(ip)
    }
}
