package main

import cc "spacemesh/internal/connectioncache"

func main() {
	var c = cc.CreateP2pConnection(1433)
	c.Open()
	c.Close()
}
