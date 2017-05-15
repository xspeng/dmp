package dmpd

type Topic struct {
	name          string
	memeryMsgChan chan *Message
}
