package event

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Hub interface {
	// Register register event handler
	Register(handler Handler) error
	// RegisterWithName register event handler by custom name
	RegisterWithName(handlerName string, handler Handler) error
	// Get event handler by event
	Get(handlerName string) ([]Handler, error)
}

type DefaultEventHub struct {
	handlerMap map[string][]Handler
}

func (p *DefaultEventHub) Get(handlerName string) (handlers []Handler, err error) {
	handlers = p.handlerMap[handlerName]
	if len(handlers) <= 0 {
		m := fmt.Sprintf("Could not found handler: %v", handlerName)
		log.Fatalf(m)
		err = errors.New(m)
	}
	return
}

func (p *DefaultEventHub) Register(handler Handler) (err error) {
	handlerType := reflect.TypeOf(handler)
	eventName := handlerType.Name()
	p.RegisterWithName(eventName, handler)
	return
}

func (p *DefaultEventHub) RegisterWithName(handlerName string, handler Handler) (err error) {
	handlers := p.handlerMap[handlerName]
	if len(handlers) == 0 {
		handlers = make([]Handler, 0)
	}
	handlers = append(handlers, handler)
	p.handlerMap[handlerName] = handlers
	return
}
