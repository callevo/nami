package nami

type Client strut {
		//nats connection handler
		mbus *nats.Conn

		_uuids string
}

// a native Asterisk AMI- server.
type Options struct {
	// Usually localhost
	Host string

	//Port to open the asterisk server
	Port int

	// Username for ARI authentication
	Username string

	// Password for ARI authentication
	Password string

	// Allow subscribe to all events in Asterisk Server
	Events string

	// Nats cluster url
	Url string
}

func NewClient() *Client {
	return &Client{}
}

func (t *Client) Connect(clientOptions Options) error {
	return nil
}

func (t *Client) Connected() bool {
	return true
}