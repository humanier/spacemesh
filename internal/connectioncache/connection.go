package connectioncache

type ConnectionStatus int

const (
	Open       ConnectionStatus = 0
	InProgress ConnectionStatus = 1
	Closed     ConnectionStatus = 2
)

type Connection interface {
	Open() error
	GetStatus() ConnectionStatus
	Close()
}
