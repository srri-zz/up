package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

var (
	host       string
	protocol   string
	ports      string
	interval   int
	timeout    int
	runTimeout int
)

func init() {
	flag.StringVar(&host, "host", "", "target host -host=x")
	flag.StringVar(&host, "h", "", "target host -h=x (shorthand)")
	flag.StringVar(&protocol, "protos", "tcp", "target protocol -proto=<tcp/udp>")
	flag.StringVar(&protocol, "pr", "tcp", "target protocol -pr=<tcp/udp> (shorthand)")
	flag.StringVar(&ports, "ports", "", "list of ports to test -ports=x,x,x")
	flag.IntVar(&interval, "interval", 1, "interval to test -interval=x")
	flag.IntVar(&interval, "i", 1, "interval to test -i=x (shorthand)")
	flag.IntVar(&timeout, "timeout", 1, "timeout -timeout=x")
	flag.IntVar(&timeout, "t", 1, "timeout -t=x (shorthand)")
	flag.Parse()
}
func main() {
	if host == "" || ports == "" {
		flag.PrintDefaults()
		return
	}
	protoList := strings.Split(protocol, ",")
	portsList := strings.Split(ports, ",")
	if len(protoList) == len(portsList) {
		for Idx, portIdx := range portsList {
			for to := 0; to <= timeout; to++ {
				conn, _ := net.DialTimeout(protoList[Idx], host+":"+portIdx, time.Second*(time.Duration(timeout)))
				if conn != nil {
					fmt.Printf("OK! - %v - %v - %v \n", host, portIdx, protoList[Idx])
					break
				} else if to == timeout {
					fmt.Printf("Fail! - %v - %v\n", host, portIdx)
				}
				time.Sleep(time.Second * (time.Duration(interval)))
			}
		}
	} else if protocol == "tcp" {
		for _, portIdx := range portsList {
			for to := 0; to <= timeout; to++ {
				conn, _ := net.DialTimeout("tcp", host+":"+portIdx, time.Second*(time.Duration(timeout)))
				if conn != nil {
					fmt.Printf("OK! - %v - %v - %v \n", host, portIdx, "tcp")
					break
				} else if to == timeout {
					fmt.Printf("Fail! - %v - %v - %v \n", host, portIdx, "tcp")
				}
				time.Sleep(time.Second * (time.Duration(interval)))
			}
		}
	} else if protocol == "udp" {
		for _, portIdx := range portsList {
			for to := 0; to <= timeout; to++ {
				conn, _ := net.DialTimeout("udp", host+":"+portIdx, time.Second*(time.Duration(timeout)))
				if conn != nil {
					fmt.Printf("OK! - %v - %v - %v \n", host, portIdx, "udp")
					break
				} else if to == timeout {
					fmt.Printf("Fail! - %v - %v - %v \n", host, portIdx, "udp")
				}
				time.Sleep(time.Second * (time.Duration(interval)))
			}
		}
	}
}
