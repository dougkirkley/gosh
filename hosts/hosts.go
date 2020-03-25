package hosts

import (
	"os"
	"log"
	"filepath"
	"github.com/kevinburke/ssh_config"
)

// ReadHosts reads hosts file
func ReadHosts() ssh_config.Hosts {
	f, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	if err != nil {
		log.Print(err.Error())
	}
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		log.Print(err.Error())
	}
	return cfg.Hosts
}