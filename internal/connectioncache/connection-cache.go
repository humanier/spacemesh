package connectioncache

type ConnectionCache interface {
	/*
		Returns existing connection or create one and store, this called should be blocked until the connection is returned. Should be thread safe.
	*/
	getConnection(ipAddress int) *Connection

	/*
		A callback function that is called whenever a remote peer establishes a connection with the local ndoe. Should ber thread safe.
	*/
	onNewRemoteConnection(remotePeer int, conn *Connection)

	/*
		Graceful shutdown, all background workers should be stopped before this method returns.
	*/
	shutdown()
}
