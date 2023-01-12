package netutil

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestHttpClient_Get(t *testing.T) {
	assert := internal.NewAssert(t, "TestHttpClient_Get")

	request := &HttpRequest{
		RawURL: "https://jsonplaceholder.typicode.com/todos/1",
		Method: "GET",
	}

	httpClient := NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(err)
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
		t.Log(err)
	}

	assert.Equal(1, todo.Id)
}

func TestHttpClent_Post(t *testing.T) {
	header := http.Header{}
	header.Add("Content-Type", "multipart/form-data")

	postData := url.Values{}
	postData.Add("userId", "1")
	postData.Add("title", "testItem")

	request := &HttpRequest{
		RawURL:   "https://jsonplaceholder.typicode.com/todos",
		Method:   "POST",
		Headers:  header,
		FormData: postData,
	}

	httpClient := NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestStructToUrlValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToUrlValues")

	type TodoQuery struct {
		Id     int `json:"id"`
		UserId int `json:"userId"`
	}
	todoQuery := TodoQuery{
		Id:     1,
		UserId: 1,
	}
	todoValues := StructToUrlValues(todoQuery)

	assert.Equal("1", todoValues.Get("id"))
	assert.Equal("1", todoValues.Get("userId"))

	request := &HttpRequest{
		RawURL:      "https://jsonplaceholder.typicode.com/todos",
		Method:      "GET",
		QueryParams: todoValues,
	}

	httpClient := NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", string(body))
}
