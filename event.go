package eventbus

type Event interface {
	GetHandlerName() string
}

type Handler interface {
	Execute(event Event) Response
}
