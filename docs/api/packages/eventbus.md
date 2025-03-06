# EventBus

EventbBus是一个事件总线，用于在应用程序中处理事件。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/eventbus/eventbus.go](https://github.com/duke-git/lancet/blob/main/eventbus/eventbus.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/eventbus"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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

## 文档

### <span id="NewEventBus">NewEventBus</span>

<p>创建EventBus实例。</p>

<b>函数签名:</b>

```go
func NewEventBus[T any]() *EventBus[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>订阅具有特定事件主题和监听函数的事件。支持异步，事件优先级，事件过滤器。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) Subscribe(topic string, listener func(eventData T), async bool, priority int, filter func(eventData T) bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>取消订阅具有特定事件主题的事件。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) Unsubscribe(topic string, listener func(eventData T))
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>发布一个带有特定事件主题和数据负载的事件。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) Publish(event Event[T])
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>清空所有事件监听器。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) ClearListeners()
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>清空特定事件监听器。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) ClearListenersByTopic(topic string)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>获取特定事件的监听器数量。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) GetListenersCount(topic string) int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>获取所有事件的监听器数量。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) GetAllListenersCount() int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>获取所有事件的topic。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) GetEvents() []string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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

<p>设置事件的错误处理函数。</p>

<b>函数签名:</b>

```go
func (eb *EventBus[T]) SetErrorHandler(handler func(err error))
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

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