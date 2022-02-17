package connectioncache

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type p2pConnection struct {
	ipAddress int
	isOpen    bool
}

func init() {
	rand.Seed((time.Now().UnixNano()))
}

func CreateP2pConnection(ipAddress int) p2pConnection {
	return p2pConnection{ipAddress, false}
}

func (c p2pConnection) Open() error {
	if c.isOpen {
		return errors.New("Connection is already open.")
	}

	// sleep random time to imitate network connection
	var timeout = rand.Intn(6) + 1
	time.Sleep(time.Second * time.Duration(timeout))

	c.isOpen = true
	fmt.Printf("- opened connection to %d in %d seconds\n", c.ipAddress, timeout)
	return nil
}

func (c p2pConnection) Close() {
	if c.isOpen {
		fmt.Printf("- cannot close connection to %d\n (not open)", c.ipAddress)
	}

	fmt.Printf("- closed connection to %d\n", c.ipAddress)
	c.isOpen = false
}
