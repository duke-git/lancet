package netutil

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestHttpGet(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := HttpGet(url, header)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	body, _ := ioutil.ReadAll(resp.Body)
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
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpPostFormData(t *testing.T) {
	apiUrl := "https://jsonplaceholder.typicode.com/todos"
	header := map[string]string{
		// "Content-Type": "application/x-www-form-urlencoded",
		"Content-Type": "multipart/form-data",
	}
	type Todo struct {
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	postData := url.Values{}
	postData.Add("userId", "1")
	postData.Add("title", "TestAddToDo")

	resp, err := HttpPost(apiUrl, header, postData, nil)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
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
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
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
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestHttpDelete(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := HttpDelete(url)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("response: ", resp.StatusCode, string(body))
}

func TestConvertMapToQueryString(t *testing.T) {
	assert := internal.NewAssert(t, "TestConvertMapToQueryString")

	var m = map[string]interface{}{
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
		t.FailNow()
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
		t.FailNow()
	}
	t.Log("response: ", toDoResp)
}
