package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	AddListener(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	RemoveListener(eventName string, handler EventHandlerInterface) error
	HasListeners(eventName string, handler EventHandlerInterface) bool
	ClearListeners() error
}
