package redis

// Connection represents a connection with redis client
type Connection interface {
	Write([]byte) error
	SetPassword(string)
	GetPassword() string

	// client should keep its subscribing channels
	Subscribe(channel string)
	UnSubscribe(channel string)
	SubsCount() int
	GetChannels() []string

}
