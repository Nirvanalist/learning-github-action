package main

import "fmt"

type Birds interface {
	Name() string
	Class() string
}

type Bird struct {
	Type string
}

func (b *Bird) Class() string {
	return b.Type
}

type Canary struct {
	Bird
	name string
}

func (c *Canary) Name() string {
	return c.name
}

func NewCanary(name string) *Canary {
	return &Canary{
		Bird: Bird{
			Type: "canary",
		},
		name: name,
	}
}

func BirdInfo(birds Birds) {
	fmt.Printf("I'm %s, I belong to %s bird class!\n", birds.Name(), birds.Class())
}

func main() {
	canary := NewCanary("cc")
	BirdInfo(canary)
}
