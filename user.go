package main

import "fmt"

type User struct {
	Name string
}

func (c *User) update(itemName string) {
	fmt.Printf("Sending email to User %s for message:  %s\n", c.Name, itemName)
}

func (c *User) getID() string {
	return c.Name
}
