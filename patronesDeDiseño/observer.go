package main

import "fmt"

// Objects can suscribed to an event
type Topic interface {
	register(observer Observer)
	roadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name string
	available bool
}

func NewItem(name string) *Item  {
	return &Item{
		name: name,
	}
}

func (i *Item)updateAvailable()  {
	fmt.Printf("Item %s is available \n",i.name)
	i.available= true
	i.broadcast()
}
func (i *Item) register(obs Observer) {
	i.observers = append(i.observers,obs)
}

func (i *Item) broadcast(){
	for _, observer := range i.observers{
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return  eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n",value, eC.id)
}

func main()  {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34ac",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.updateAvailable()

}