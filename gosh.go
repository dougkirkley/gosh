package main

import (
	"flag"
	"fmt"
	"strings"
	_"time"
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
	puppetArgs    = "--onetime --verbose --no-daemonize --ignorecache --no-usecacheonfailure --no-splay"
	puppetCmd     = "cd /tmp; sudo /usr/bin/puppet agent"
	defaultCmd    = fmt.Sprintf("%s %s", puppetCmd, puppetArgs)
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
	Cmd := flag.String("c", defaultCmd, "Command to Run on Remote Nodes")
	Hosts := flag.String("H", strings.Join(hosts, ","), "Hosts to Run Command on")
	Quiet := flag.Bool("q", false, "Hide Output From Commands")
	flag.Parse()
	if cmp.Equal(hosts, Hosts) {
		fmt.Println("No Host entered. Running on all detected servers")
		fmt.Print("Running on all servers")
	}
	for _, hostname := range hosts {
		go func(hostname string) {
			c := &Connection{
				Host:    hostname,
				Command: *Cmd,
				Quiet:   *Quiet,
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
