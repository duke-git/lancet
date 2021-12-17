package netutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestHttpGet(t *testing.T) {
	_, e := HttpGet("", nil)
	if e == nil {
		t.FailNow()
	}

	url := "https://gutendex.com/books?"
	queryParams := make(map[string]interface{})
	queryParams["ids"] = "1"

	resp, err := HttpGet(url, nil, queryParams)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", resp.StatusCode, string(body))

}

func TestHttpPost(t *testing.T) {
	url := "http://public-api-v1.aspirantzhang.com/users"
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	user := User{
		"test",
		"test@test.com",
	}
	bodyParams, _ := json.Marshal(user)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HttpPost(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", resp.StatusCode, string(body))
}

func TestHttpPut(t *testing.T) {
	url := "http://public-api-v1.aspirantzhang.com/users/10420"
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	user := User{
		"test",
		"test@test.com",
	}
	bodyParams, _ := json.Marshal(user)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HttpPut(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", resp.StatusCode, string(body))
}

func TestHttpDelete(t *testing.T) {
	url := "http://public-api-v1.aspirantzhang.com/users/10420"
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
		utils.LogFailedTestInfo(t, "ConvertMapToQueryString", m, expected, r)
		t.FailNow()
	}
}

func TestParseResponse(t *testing.T) {
	url := "http://public-api-v1.aspirantzhang.com/users"
	type User struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	type UserResp struct {
		Data []User `json:"data"`
	}
	resp, err := HttpGet(url)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	userResp := &UserResp{}
	err = ParseResponse(resp, userResp)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	fmt.Println(userResp.Data)
}
