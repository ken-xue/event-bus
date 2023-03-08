package eventbus

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
	t.Log(eventBus)
	Register(PipelineOnlineEventHandler{})
	RegisterWithName(PipelineOnlineEvent{}.GetHandlerName(), PipelineOnlineEventHandler{})
	Publish(PipelineOnlineEvent{"test"})
	PublishAll(PipelineOnlineEvent{"test"})
	AsyncPublish(PipelineOnlineEvent{"test"})
	AsyncPublishAll(PipelineOnlineEvent{"test"})
}
