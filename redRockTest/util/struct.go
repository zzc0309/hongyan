package util

type Message struct {
	Username string
	Message string
	ChildMessage *[]Message
}
