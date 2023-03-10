package eventbus

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

func NewDefaultEventHub() *DefaultEventHub {
	return &DefaultEventHub{handlerMap: make(map[string][]Handler)}
}

func (p *DefaultEventHub) Get(handlerName string) (handlers []Handler, err error) {
	handlers = p.handlerMap[handlerName]
	if len(handlers) <= 0 {
		err = errors.New(fmt.Sprintf("Could not found handler: %v", handlerName))
	}
	return
}

func (p *DefaultEventHub) Register(handler Handler) (err error) {
	handlerType := reflect.TypeOf(handler)
	handlerName := handlerType.Name()
	log.Println("register handler : ", handlerName)
	return p.RegisterWithName(handlerName, handler)
}

func (p *DefaultEventHub) RegisterWithName(handlerName string, handler Handler) (err error) {
	handlers := p.handlerMap[handlerName]
	if len(handlers) == 0 {
		handlers = make([]Handler, 0)
	}
	handlers = append(handlers, handler)
	log.Println("register handler : ", handlerName)
	p.handlerMap[handlerName] = handlers
	return
}
