package eventbus

import (
	"sync"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestEventBus_Subscribe(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_Subscribe")

	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {
		assert.Equal(1, eventData)
	}, false, 0, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
}

func TestEventBus_Unsubscribe(t *testing.T) {

	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_Unsubscribe")

	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	eb.Subscribe("event1", listener, false, 0, nil)
	eb.Unsubscribe("event1", listener)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	assert.Equal(0, receivedData)
}

func TestEventBus_Subscribe_WithFilter(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_Subscribe_WithFilter")

	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	filter := func(eventData int) bool {
		return eventData == 1
	}

	eb.Subscribe("event1", listener, false, 0, filter)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	eb.Publish(Event[int]{Topic: "event1", Payload: 2})

	assert.Equal(1, receivedData)
}

func TestEventBus_Subscribe_WithPriority(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_Subscribe_WithPriority")

	eb := NewEventBus[int]()

	var receivedData []int
	listener1 := func(eventData int) {
		receivedData = append(receivedData, 1)
	}

	listener2 := func(eventData int) {
		receivedData = append(receivedData, 2)
	}

	eb.Subscribe("event1", listener1, false, 1, nil)
	eb.Subscribe("event1", listener2, false, 2, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	assert.Equal([]int{2, 1}, receivedData)
}

func TestEventBus_Subscribe_Async(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_Subscribe_Async")

	eb := NewEventBus[string]()

	var wg sync.WaitGroup
	wg.Add(1)

	eb.Subscribe("event1", func(eventData string) {
		time.Sleep(100 * time.Millisecond)
		assert.Equal("hello", eventData)
		wg.Done()
	}, true, 1, nil)

	eb.Publish(Event[string]{Topic: "event1", Payload: "hello"})

	wg.Wait()
}

func TestEventBus_ErrorHandler(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_ErrorHandler")

	eb := NewEventBus[string]()

	eb.SetErrorHandler(func(err error) {
		assert.Equal("error", err.Error())
	})

	eb.Subscribe("event1", func(eventData string) {
		panic("error")
	}, false, 0, nil)

	eb.Publish(Event[string]{Topic: "event1", Payload: "hello"})
}

func TestEventBus_ClearListeners(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_ClearListeners")

	eb := NewEventBus[int]()

	receivedData1 := 0
	receivedData2 := 0

	eb.Subscribe("event1", func(eventData int) {
		receivedData1 = eventData
	}, false, 0, nil)

	eb.Subscribe("event2", func(eventData int) {
		receivedData2 = eventData
	}, false, 0, nil)

	eb.ClearListeners()

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	eb.Publish(Event[int]{Topic: "event1", Payload: 2})

	assert.Equal(0, receivedData1)
	assert.Equal(0, receivedData2)
}

func TestEventBus_ClearListenersByTopic(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_ClearListenersByTopic")

	eb := NewEventBus[int]()

	receivedData1 := 0
	receivedData2 := 0

	eb.Subscribe("event1", func(eventData int) {
		receivedData1 = eventData
	}, false, 0, nil)

	eb.Subscribe("event2", func(eventData int) {
		receivedData2 = eventData
	}, false, 0, nil)

	eb.ClearListenersByTopic("event1")

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	eb.Publish(Event[int]{Topic: "event2", Payload: 2})

	assert.Equal(0, receivedData1)
	assert.Equal(2, receivedData2)
}

func TestEventBus_GetListenersCount(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_GetListenersCount")

	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	assert.Equal(2, eb.GetListenersCount("event1"))
	assert.Equal(1, eb.GetListenersCount("event2"))
}

func TestEventBus_GetAllListenersCount(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_GetAllListenersCount")

	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	assert.Equal(3, eb.GetAllListenersCount())
}

func TestEventBus_GetEvents(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEventBus_GetEvents")

	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	events := eb.GetEvents()

	assert.Equal(2, len(events))
	assert.Equal("event1", events[0])
	assert.Equal("event2", events[1])
}
