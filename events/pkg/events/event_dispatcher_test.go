package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {
	// do nothing
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event          TestEvent
	anotherEvent   TestEvent
	handler        TestEventHandler
	anotherHandler TestEventHandler
	otherHandler   TestEventHandler
	dispatcher     *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.anotherEvent = TestEvent{Name: "another", Payload: "another"}
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.anotherHandler = TestEventHandler{
		ID: 2,
	}
	suite.otherHandler = TestEventHandler{
		ID: 3,
	}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.AddListener(suite.event.GetName(), &suite.anotherHandler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 2)

	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.anotherHandler, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterTwice() {
	err := suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Equal(suite.T(), ErrHandlerAlreadyRegistered, err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.AddListener(suite.event.GetName(), &suite.anotherHandler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 2)

	err = suite.dispatcher.AddListener(suite.anotherEvent.GetName(), &suite.otherHandler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.anotherEvent.GetName()], 1)

	suite.dispatcher.ClearListeners()
	suite.Equal(0, len(suite.dispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_HasListeners() {
	err := suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.AddListener(suite.event.GetName(), &suite.anotherHandler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 2)

	assert.True(suite.T(), suite.dispatcher.HasListeners(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.dispatcher.HasListeners(suite.event.GetName(), &suite.anotherHandler))
	assert.False(suite.T(), suite.dispatcher.HasListeners(suite.event.GetName(), &suite.otherHandler))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.dispatcher.AddListener(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.AddListener(suite.event.GetName(), &suite.anotherHandler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 2)

	suite.dispatcher.RemoveListener(suite.event.GetName(), &suite.handler)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 1)
	assert.Equal(suite.T(), &suite.anotherHandler, suite.dispatcher.handlers[suite.event.GetName()][0])

	suite.dispatcher.RemoveListener(suite.event.GetName(), &suite.anotherHandler)
	assert.Len(suite.T(), suite.dispatcher.handlers[suite.event.GetName()], 0)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &suite.event)
	suite.dispatcher.AddListener(suite.event.GetName(), eventHandler)
	suite.dispatcher.Dispatch(&suite.event)
	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
