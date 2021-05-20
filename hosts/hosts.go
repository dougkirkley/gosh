package hosts

import (
	"github.com/kevinburke/ssh_config"
)

// GetIPFromHostname gets the ip from the hostname in your ssh config
func GetIPFromHostname(host string) string {
	ip := ssh_config.Get(host, "Hostname")
	return ip
}