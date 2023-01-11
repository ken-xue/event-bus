package event

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Hub interface {
	// Register register event handler
	Register(eventHandler Handler) error
	// Get event handler by event
	Get(event Event) (Handler, error)
}

type DefaultEventHub struct {
	handlerMap map[string]Handler
}

func (p DefaultEventHub) Get(event Event) (handler Handler, err error) {
	handler, exist := p.handlerMap[event.GetHandlerName()]
	if !exist {
		m := fmt.Sprintf("Could not found handler: %v", event)
		log.Fatalf(m)
		err = errors.New(m)
	}
	return
}

func (p DefaultEventHub) Register(handler Handler) (err error) {
	handlerType := reflect.TypeOf(handler)
	eventName := handlerType.Name()
	_, exists := p.handlerMap[eventName]
	if exists {
		m := fmt.Sprintf("Event handler already exists : %v", eventName)
		log.Fatalf(m)
		err = errors.New(m)
		return
	}
	p.handlerMap[eventName] = handler
	return
}
