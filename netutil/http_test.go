package netutil

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestHttpGet(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := HttpGet(url, header)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpPost(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, "TestAddToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := HttpPost(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpPostFormData(t *testing.T) {
	apiUrl := "https://jsonplaceholder.typicode.com/todos"
	header := map[string]string{
		// "Content-Type": "application/x-www-form-urlencoded",
		"Content-Type": "multipart/form-data",
	}

	postData := url.Values{}
	postData.Add("userId", "1")
	postData.Add("title", "TestToDo")

	// postData := make(map[string]string)
	// postData["userId"] = "1"
	// postData["title"] = "title"

	resp, err := HttpPost(apiUrl, header, postData, nil)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpPut(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, 1, "TestPutToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := HttpPut(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpPatch(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, 1, "TestPatchToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := HttpPatch(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpDelete(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := HttpDelete(url)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestConvertMapToQueryString(t *testing.T) {
	assert := internal.NewAssert(t, "TestConvertMapToQueryString")

	var m = map[string]any{
		"c": 3,
		"a": 1,
		"b": 2,
	}
	assert.Equal("a=1&b=2&c=3", ConvertMapToQueryString(m))
}

func TestParseResponse(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := HttpGet(url, header)
	if err != nil {
		log.Fatal(err)
	}

	type Todo struct {
		Id        int    `json:"id"`
		UserId    int    `json:"userId"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	toDoResp := &Todo{}
	err = ParseHttpResponse(resp, toDoResp)
	if err != nil {
		log.Fatal(err)
	}
	t.Log("response: ", toDoResp)
}

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
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Name   string `json:"name,omitempty"`
	}
	item1 := TodoQuery{
		Id:     1,
		UserId: 123,
		Name:   "",
	}
	todoValues, err := StructToUrlValues(item1)
	if err != nil {
		t.Errorf("params is invalid: %v", err)
	}

	assert.Equal("1", todoValues.Get("id"))
	assert.Equal("123", todoValues.Get("userId"))
	assert.Equal("", todoValues.Get("name"))

	item2 := TodoQuery{
		Id:     2,
		UserId: 456,
	}
	queryValues2, _ := StructToUrlValues(item2)

	assert.Equal("2", queryValues2.Get("id"))
	assert.Equal("456", queryValues2.Get("userId"))
	assert.Equal("", queryValues2.Get("name"))
}
