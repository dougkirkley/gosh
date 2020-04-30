package main

import (
	"flag"
	"fmt"
	"strings"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"net"

	"github.com/google/go-cmp/cmp"
	"github.com/dougkirkley/gosh/connection"
	"github.com/dougkirkley/gosh/hosts"
)

var (
        Hosts         = hosts.ReadHosts()
	results       = make(chan []string, 150)
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f)
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	Cmd := flag.String("c", "", "Command to Run on Remote Nodes")
	Hosts := flag.String("H", strings.Join(hosts, ","), "Hosts to Run Command on")
	flag.Parse()
	for _, hostname := range hosts {
		go func(hostname string) {
			c := &Connection{
				Host:    hostname,
				Command: *Cmd,
				Timeout: *Timeout,
			}
			results <- c.ExecuteCmd()
		}(hostname)
	}
	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			for _, line := range res {
				fmt.Println(line)
			}
		}
	}
	close(results)
}
