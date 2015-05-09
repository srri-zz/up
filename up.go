package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	host       string
	protocol   string
	port       string
	interval   int
	timeout    int
	runTimeout int
	verbose    string
)

func init() {
	flag.StringVar(&host, "host", "", "target host -host=x")
	flag.StringVar(&host, "h", "", "target host -h=x (shorthand)")
	flag.StringVar(&protocol, "proto", "tcp", "target protocol -proto=<tcp/udp>")
	flag.StringVar(&protocol, "pr", "tcp", "target protocol -pr=<tcp/udp> (shorthand)")
	flag.StringVar(&port, "port", "", "target port -port=x")
	flag.StringVar(&port, "p", "", "target port -p=x (shorthand)")
	flag.StringVar(&verbose, "verbose", "false", "target port -verbose=x")
	flag.StringVar(&verbose, "v", "false", "target port -v=x (shorthand)")
	flag.IntVar(&interval, "interval", 1, "interval to test -interval=x")
	flag.IntVar(&interval, "i", 1, "interval to test -i=x (shorthand)")
	flag.IntVar(&timeout, "timeout", 1, "timeout -timeout=x")
	flag.IntVar(&timeout, "t", 1, "timeout -t=x (shorthand)")
	flag.Parse()
}
func main() {
	if host == "" || port == "" {
		flag.PrintDefaults()
		return
	}
	if verbose == "true" {
		fmt.Println(host)
		fmt.Println(protocol)
		fmt.Println(port)
		fmt.Println(interval)
		fmt.Println(timeout)
	}
	for {
		conn, _ := net.DialTimeout(protocol, host+":"+port, time.Second*(time.Duration(timeout)))
		if verbose == "true" {
			fmt.Println(conn)
		}
		if conn != nil {
			fmt.Print(host)
			return
		}
		time.Sleep(time.Second * (time.Duration(interval)))
	}
}
