package eventbus

import (
	"fmt"
	"sync"
	"time"
)

func ExampleEventBus_Subscribe() {
	eb := NewEventBus[string]()
	eb.Subscribe("event1", func(eventData string) {
		fmt.Println(eventData)
	}, false, 0, nil)

	eb.Publish(Event[string]{Topic: "event1", Payload: "hello"})

	// Output:
	// hello
}

func ExampleEventBus_Unsubscribe() {
	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	eb.Subscribe("event1", listener, false, 0, nil)
	eb.Unsubscribe("event1", listener)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	fmt.Println(receivedData)

	// Output:
	// 0
}

func ExampleEventBus_Subscribe_withFilter() {
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

	fmt.Println(receivedData)

	// Output:
	// 1
}

func ExampleEventBus_Subscribe_withPriority() {
	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {
		fmt.Println(eventData)
	}, false, 0, nil)

	eb.Subscribe("event1", func(eventData int) {
		fmt.Println(eventData)
	}, false, 1, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	// Output:
	// 1
	// 1
}

func ExampleEventBus_Subscribe_async() {
	eb := NewEventBus[int]()

	var wg sync.WaitGroup
	wg.Add(1)

	eb.Subscribe("event1", func(eventData int) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(eventData)
		wg.Done()
	}, true, 1, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	wg.Wait()

	// Output:
	// 1
}

func ExampleEventBus_Publish() {
	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {
		fmt.Println(eventData)
	}, false, 0, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	// Output:
	// 1
}

func ExampleEventBus() {
	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	eb.Subscribe("event1", listener, false, 0, nil)
	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	fmt.Println(receivedData)

	// Output:
	// 1
}

func ExampleEventBus_ClearListeners() {
	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	eb.Subscribe("event1", listener, false, 0, nil)
	eb.Subscribe("event2", listener, false, 0, nil)

	eb.ClearListeners()

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	eb.Publish(Event[int]{Topic: "event2", Payload: 2})

	fmt.Println(receivedData)

	// Output:
	// 0
}

func ExampleEventBus_ClearListenersByTopic() {
	eb := NewEventBus[int]()

	receivedData := 0
	listener := func(eventData int) {
		receivedData = eventData
	}

	eb.Subscribe("event1", listener, false, 0, nil)
	eb.Subscribe("event2", listener, false, 0, nil)

	eb.ClearListenersByTopic("event1")

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})
	eb.Publish(Event[int]{Topic: "event2", Payload: 2})

	fmt.Println(receivedData)

	// Output:
	// 2
}

func ExampleEventBus_GetListenersCount() {
	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	count := eb.GetListenersCount("event1")

	fmt.Println(count)

	// Output:
	// 1
}

func ExampleEventBus_SetErrorHandler() {
	eb := NewEventBus[int]()

	eb.SetErrorHandler(func(err error) {
		fmt.Println(err)
	})

	eb.Subscribe("event1", func(eventData int) {
		panic("error")
	}, false, 0, nil)

	eb.Publish(Event[int]{Topic: "event1", Payload: 1})

	// Output:
	// error
}

func ExampleEventBus_GetAllListenersCount() {

	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	count := eb.GetAllListenersCount()

	fmt.Println(count)

	// Output:
	// 2
}

func ExampleEventBus_GetEvents() {
	eb := NewEventBus[int]()

	eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
	eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

	events := eb.GetEvents()

	for _, event := range events {
		fmt.Println(event)
	}

	// Output:
	// event1
	// event2
}
