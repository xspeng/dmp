package dmpd

type Message struct {
	messageID int64
	body      []byte
	timestamp int64
}

func NewMessage() *Message {
	return &Message{}
}
