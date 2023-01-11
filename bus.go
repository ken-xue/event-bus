package event

import (
	"log"
)

type Bus interface {
	publish(event Event) Response
	asyncPublish(event Event)
}

type DefaultEventBus struct {
	Hub
}

func (p DefaultEventBus) Publish(event Event) (response Response) {
	handler, err := p.Hub.Get(event)
	if err != nil {
		log.Fatal("publish error :", err)
		return
	}
	response = handler.Execute(event)
	return
}

func (p DefaultEventBus) AsyncPublish(event Event) {
	go func() {
		p.Publish(event)
	}()
}

var EventBus = &DefaultEventBus{
	Hub: &DefaultEventHub{
		handlerMap: make(map[string]Handler),
	},
}
