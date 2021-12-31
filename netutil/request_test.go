package netutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Println("response: ", resp.StatusCode, string(body))
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
	fmt.Println("response: ", resp.StatusCode, string(body))
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
	fmt.Println("response: ", resp.StatusCode, string(body))
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
	fmt.Println("response: ", resp.StatusCode, string(body))
}

func TestHttpDelete(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := HttpDelete(url)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", resp.StatusCode, string(body))
}

func TestConvertMapToQueryString(t *testing.T) {
	var m = map[string]interface{}{
		"c": 3,
		"a": 1,
		"b": 2,
	}

	expected := "a=1&b=2&c=3"
	r := ConvertMapToQueryString(m)
	if r != expected {
		internal.LogFailedTestInfo(t, "ConvertMapToQueryString", m, expected, r)
		t.FailNow()
	}
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
	fmt.Println("response: ", toDoResp)
}
