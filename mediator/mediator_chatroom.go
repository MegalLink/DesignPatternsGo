package mediator

import (
	"fmt"
)

type Person struct {
	Name    string
	Room    *ChatRoom // this line is important for mediator pattern bacause chatRoom is the mediator
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) ReceiveMessage(sender, message string) {
	s := fmt.Sprintf("%s: %s\n", sender, message)
	fmt.Printf("[%s's chat session]: %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) SendChatRoomMessage(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

// the mediator is the Chatroom because we have people in the chat but they are not aware of each other, chat room makes the connection between them
// they dont have pointer to one or other
type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.ReceiveMessage(source, message)
		}
	}
}

func (c *ChatRoom) Message(source, destination, message string) {
	for _, p := range c.people {
		if p.Name == destination {
			p.ReceiveMessage(source, message)
		}
	}
}

func (c *ChatRoom) JoinRoom(p *Person) {
	joinMsg := p.Name + " joined the chat"
	c.Broadcast("Room", joinMsg)
	p.Room = c
	c.people = append(c.people, p)
}
