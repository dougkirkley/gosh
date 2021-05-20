package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/dougkirkley/gosh/connection"
	"github.com/dougkirkley/gosh/hosts"
)

var (
	results  = make(chan []string, 150)
)

func main() {
	Cmd := flag.String("c", "", "Command to Run on Remote Nodes")
	Hosts := flag.String("H", "", "Hosts to Run Command on")
	Password := flag.String("p", "", "Password to use for ssh")
	flag.Parse()
	HostList := strings.Split(*Hosts, ",")
	for _, hostname := range HostList {
		hostname = hosts.GetIPFromHostname(hostname)
		go func(hostname string) {
			c := &connection.Connection{
				Host:    hostname,
				Command: *Cmd,
				Password: *Password,
			}
			results <- c.ExecuteCmd()
		}(hostname)
	}
	for i := 0; i < len(HostList); i++ {
		select {
		case res := <-results:
			for _, line := range res {
				fmt.Println(line)
			}
		}
	}
	close(results)
}
