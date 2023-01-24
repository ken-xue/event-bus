# Event-Bus

> A simple event-driven implementation for business decoupling


# QuickStart

```bash
go get github/ken-xue/event-bus
```

```go
package event

import (
	"fmt"
	"reflect"
	"testing"
)

type PipelineOnlineEvent struct {
	name string
}

func (p PipelineOnlineEvent) GetHandlerName() string {
	return reflect.TypeOf(PipelineOnlineEventHandler{}).Name()
}

type PipelineOnlineEventHandler struct {
}

func (p PipelineOnlineEventHandler) Execute(event Event) (re Response) {
	e := event.(PipelineOnlineEvent)
	fmt.Println(e.name)
	fmt.Println("to do something")
	return
}

func TestEvent(t *testing.T) {
	t.Log(EventBus)
	EventBus.Register(PipelineOnlineEventHandler{})
	EventBus.Publish(PipelineOnlineEvent{"test"})
}
```