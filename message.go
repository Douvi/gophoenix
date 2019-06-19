package gophoenix

// Message is a message sent or received via the Transport from the channel.
type Message struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     int64       `json:"ref"`
}

type RefMessage struct {
	UserID  string `json:"user_id"`
	Topic   string `json:"topic"`
	Ref     int64  `json:"ref"`
	JoinRef int64  `json:"join_ref"`
}

type PlayloadIn struct {
	Ref     RefMessage  `json:"ref"`
	Payload interface{} `json:"payload"`
}

type PlayloadOut struct {
	Ref     RefMessage  `json:"ref"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}
