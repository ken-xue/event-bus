package eventbus

import (
	"log"
)

type Bus interface {
	Publish(event Event) Response
	PublishAll(event Event) []Response
	AsyncPublish(event Event)
	AsyncPublishAll(event Event)
}

type DefaultEventBus struct {
	Hub
}

func (p *DefaultEventBus) Publish(event Event) (response Response) {
	handlers, err := p.Hub.Get(event.GetHandlerName())
	if err != nil {
		log.Println("publish error :", err)
		return
	}
	if len(handlers) <= 0 {
		log.Println("not found handler for event :", event)
		return
	}
	response = handlers[0].Execute(event)
	return
}

func (p *DefaultEventBus) PublishAll(event Event) (responses []Response) {
	handlers, err := p.Hub.Get(event.GetHandlerName())
	if err != nil {
		log.Println("publish error :", err)
		return
	}
	if len(handlers) <= 0 {
		log.Println("not found handler for event :", event)
		return
	}
	responses = make([]Response, len(handlers))
	for i, handler := range handlers {
		responses[i] = handler.Execute(event)
	}
	return
}

func (p *DefaultEventBus) AsyncPublish(event Event) {
	go func() {
		p.Publish(event)
	}()
}

func (p *DefaultEventBus) AsyncPublishAll(event Event) {
	go func() {
		p.PublishAll(event)
	}()
}

func NewEventBus(hub Hub) *DefaultEventBus {
	return &DefaultEventBus{Hub: hub}
}
