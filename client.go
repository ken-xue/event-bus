package eventbus

var (
	hub      Hub = NewDefaultEventHub()
	eventBus Bus = NewEventBus(hub)
)

func SetClient(bus Bus) {
	eventBus = bus
}

func Publish(event Event) (response Response) {
	return eventBus.Publish(event)
}

func PublishAll(event Event) (responses []Response) {
	return eventBus.PublishAll(event)
}

func AsyncPublish(event Event) {
	eventBus.AsyncPublish(event)
}

func AsyncPublishAll(event Event) {
	eventBus.AsyncPublishAll(event)
}

func Register(handler Handler) (err error) {
	return hub.Register(handler)
}

func RegisterWithName(handlerName string, handler Handler) (err error) {
	return hub.RegisterWithName(handlerName, handler)
}
