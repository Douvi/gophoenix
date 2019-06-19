package gophoenix

// Message is a message sent or received via the Transport from the channel.
type Message struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     int64       `json:"ref"`
}

type PlayloadIn struct {
	Ref     interface{} `json:"ref"`
	Payload interface{} `json:"payload"`
}

type PlayloadOut struct {
	Ref     interface{} `json:"ref"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}
