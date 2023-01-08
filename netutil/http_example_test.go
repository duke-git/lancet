package netutil

import "fmt"

func ExampleHttpClient_SendRequest() {
	request := &HttpRequest{
		RawURL: "https://jsonplaceholder.typicode.com/todos/1",
		Method: "GET",
	}

	httpClient := NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil || resp.StatusCode != 200 {
		return
	}

	type Todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var todo Todo
	err = httpClient.DecodeResponse(resp, &todo)
	if err != nil {
		return
	}

	fmt.Println(todo.Id)

	// Output:
	// 1
}

func ExampleHttpClient_DecodeResponse() {
	request := &HttpRequest{
		RawURL: "https://jsonplaceholder.typicode.com/todos/1",
		Method: "GET",
	}

	httpClient := NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil || resp.StatusCode != 200 {
		return
	}

	type Todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var todo Todo
	err = httpClient.DecodeResponse(resp, &todo)
	if err != nil {
		return
	}

	fmt.Println(todo.Id)

	// Output:
	// 1
}

func ExampleStructToUrlValues() {
	type TodoQuery struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	todoQuery := TodoQuery{
		Id:   1,
		Name: "Test",
	}
	todoValues := StructToUrlValues(todoQuery)

	fmt.Println(todoValues.Get("id"))
	fmt.Println(todoValues.Get("name"))

	// Output:
	// 1
	// Test
}

func ExampleConvertMapToQueryString() {
	var m = map[string]any{
		"c": 3,
		"a": 1,
		"b": 2,
	}

	qs := ConvertMapToQueryString(m)

	fmt.Println(qs)

	// Output:
	// a=1&b=2&c=3
}
