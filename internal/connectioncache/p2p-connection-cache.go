package connectioncache

import "sync"

type P2pConnectionCache struct {
	pool                      map[int]*Connection
	poolMutex                 sync.Mutex
	shutdownMutex				synx.Mutex
	getWorkers                sync.WaitGroup
	shutdownSequenceInitiated bool
}

/*
	Returns existing connection or create one and store, this called should be blocked until the connection is returned. Should be thread safe.
*/
func (c P2pConnectionCache) getConnection(ipAddress int) *Connection {
	c.shutdownMutex.Lock()
	if c.shutdownSequenceInitiated {
		return nil
	}
	else {
		c.getWorkers.Add(1)
	}

	connChannel := make(chan *Connection)
	
	go func() {	
		defer c.getWorkers.Done()
		c.poolMutex.Lock()
		

		conn, exists := c.pool[ipAddress]
		if !exists {
			conn = P2pConnection(ipAddress)
			c.pool[ipAddress] = &conn
		}

		c.poolMutex.Unlock()

		// conn.Open() 
		// or
		// wait until conn.GetStatus() is Open

		connChannel <- conn		
	}

	return <- connChannel
} 

/*
	A callback function that is called whenever a remote peer establishes a connection with the local node. Should ber thread safe.
*/
func (c P2pConnectionCache) onNewRemoteConnection(remotePeer int, conn *Connection) {

}

/*
	Graceful shutdown, all background workers should be stopped before this method returns.
*/
func (c P2pConnectionCache) shutdown() {
	c.shutdownSequenceInitiated = true

	// wait until getWorkers WaitGroup is drained
	c.getWorkers.Wait()

	// iterate on the pool closing connections one by one	
}
