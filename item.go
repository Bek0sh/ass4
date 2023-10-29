package main

import "fmt"

type ChatRoom struct {
	observerList []Observer
	message      string
}

func newChat(message string) *ChatRoom {
	return &ChatRoom{
		message: message,
	}
}

func (i *ChatRoom) updateAvailability() {
	fmt.Printf("New message: %s \n", i.message)
	i.notifyAll()
}

func (i *ChatRoom) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *ChatRoom) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *ChatRoom) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.message)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
