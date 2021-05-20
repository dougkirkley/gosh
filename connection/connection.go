package connection

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

// Connection for setting up an ssh session for remote commands
type Connection struct {
	Host     string
	Password string
	Command  string
}

// PublicKeyFile reads in users ssh key
func (c *Connection) PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

// MakeSSHConfig Gets ssh key from home dir
func (c *Connection) MakeSSHConfig() *ssh.ClientConfig {
	user := os.Getenv("USER")
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			c.PublicKeyFile(fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))),
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

// Connect makes ssh client for host
func (c *Connection) Connect() (*ssh.Session, io.Reader, io.Reader) {
	config := c.MakeSSHConfig()
	client, err := ssh.Dial("tcp", c.Host+":22", config)
	if err != nil {
		log.Fatal(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		log.Fatalf("request for pseudo terminal failed: %v", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Unable to setup stdout for session: %v", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("Unable to setup stderr for session: %v", err)
	}

	return session, stdout, stderr
}

// ExecuteCmd runs commands on remote host
func (c *Connection) ExecuteCmd() []string {
	var output []string
	session, stdout, stderr := c.Connect()
	defer session.Close()
	session.Run(c.Command)
	scanner := bufio.NewScanner(stdout)
	go io.Copy(os.Stderr, stderr)

	for scanner.Scan() {
		if scanner.Text() != "" {
			output = append(output, c.Host+": "+scanner.Text())
		} else {
			break
		}
		if scanner.Err() != nil {
			log.Printf("error: %s\n", scanner.Err())
		}
	}
	return output
}
