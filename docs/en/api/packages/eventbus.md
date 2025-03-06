# EventBus

EventBus is an event bus used for handling events within an application.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/eventbus/eventbus.go](https://github.com/duke-git/lancet/blob/main/eventbus/eventbus.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/eventbus"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [NewEventBus](#NewEventBus)
-   [Subscribe](#Subscribe)
-   [Unsubscribe](#Unsubscribe)
-   [Publish](#Publish)
-   [ClearListeners](#ClearListeners)
-   [ClearListenersByTopic](#ClearListenersByTopic)
-   [GetListenersCount](#GetListenersCount)
-   [GetAllListenersCount](#GetAllListenersCount)
-   [GetEvents](#GetEvents)
-   [SetErrorHandler](#SetErrorHandler)


<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="NewEventBus">NewEventBus</span>

<p>Create an EventBus instance.</p>

<b>Signature:</b>

```go
func NewEventBus[T any]() *EventBus[T]
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

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
```

### <span id="Subscribe">Subscribe</span>

<p>Subscribes to an event with a specific event topic and listener function.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) Subscribe(topic string, listener func(eventData T), async bool, priority int, filter func(eventData T) bool)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

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
```

### <span id="Unsubscribe">Unsubscribe</span>

<p>Unsubscribes from an event with a specific event topic and listener function.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) Unsubscribe(topic string, listener func(eventData T))
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

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
```

### <span id="Publish">Publish</span>

<p>Publishes an event with a specific event topic and data payload.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) Publish(event Event[T])
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

    eb.Subscribe("event1", func(eventData int) {
        fmt.Println(eventData)
    }, false, 0, nil)

    eb.Publish(Event[int]{Topic: "event1", Payload: 1})

    // Output:
    // 1
}
```

### <span id="ClearListeners">ClearListeners</span>

<p>Clears all the listeners.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) ClearListeners()
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

    receivedData := 0
    listener := func(eventData int) {
        receivedData = eventData
    }

    eb.Subscribe("event1", listener, false, 0, nil)
    eb.ClearListeners()

    eb.Publish(Event[int]{Topic: "event1", Payload: 1})

    fmt.Println(receivedData)

    // Output:
    // 0
}
```

### <span id="ClearListenersByTopic">ClearListenersByTopic</span>

<p>Clears all the listeners by topic.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) ClearListenersByTopic(topic string)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

    receivedData := 0
    listener := func(eventData int) {
        receivedData = eventData
    }

    eb.Subscribe("event1", listener, false, 0, nil)
    eb.ClearListenersByTopic("event1")

    eb.Publish(Event[int]{Topic: "event1", Payload: 1})

    fmt.Println(receivedData)

    // Output:
    // 0
}
```

### <span id="GetListenersCount">GetListenersCount</span>

<p>Returns the number of listeners for a specific event topic.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) GetListenersCount(topic string) int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

    eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
    eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)

    count := eb.GetListenersCount("event1")

    fmt.Println(count)

    // Output:
    // 2
}
```

### <span id="GetAllListenersCount">GetAllListenersCount</span>

<p>Returns the total number of all listeners.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) GetAllListenersCount() int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

    eb.Subscribe("event1", func(eventData int) {}, false, 0, nil)
    eb.Subscribe("event2", func(eventData int) {}, false, 0, nil)

    count := eb.GetAllListenersCount()

    fmt.Println(count)

    // Output:
    // 2
}
```

### <span id="GetEvents">GetEvents</span>

<p>Returns all the events topics.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) GetEvents() []string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

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
```

### <span id="SetErrorHandler">SetErrorHandler</span>

<p>Sets the error handler function.</p>

<b>Signature:</b>

```go
func (eb *EventBus[T]) SetErrorHandler(handler func(err error))
```

<b>Example:<span style="float:right;display:inline-block;">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/eventbus"
)

func main() {
    eb := eventbus.NewEventBus[int]()

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
```