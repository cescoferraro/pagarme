package infra

import (
	"net/http/httptest"
	"testing"

	"fmt"

	"io/ioutil"
	"net/http"

	"encoding/json"
	"io"
	"reflect"
	"strings"

	"github.com/stretchr/testify/assert"
)

//TestServer is the global server var
var TestServer *httptest.Server

//TableTest is an object
type TableTest struct {
	Method       string
	Path         string
	Jwt          string
	Body         interface{}
	Expect       interface{}
	BodyContains string
	Status       int
	Name         string
	Description  string
}

// Spin runs the test
func (test TableTest) Spin(t *testing.T) string {
	return string(test.innnerSpin(t))
}

func (test TableTest) innnerSpin(t *testing.T) []byte {
	url := TestServer.URL + test.Path
	b, err := json.Marshal(test.Body)
	assert.NoError(t, err)
	r, err := http.NewRequest(test.Method, url, strings.NewReader(string(b)))
	assert.NoError(t, err)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", test.Jwt))
	r.Header.Set("CESCO", fmt.Sprintf("CESCO  won %v", test.Jwt))
	response, err := http.DefaultClient.Do(r)
	assert.NoError(t, err)
	actualBody, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(actualBody), test.BodyContains, "body")
	assert.Equal(t, test.Status, response.StatusCode, "status code")
	return actualBody
}

// DoubleSpin run multiple teste
func (test TableTest) DoubleSpin(t *testing.T) interface{} {
	actualBody := test.innnerSpin(t)
	thetype := reflect.TypeOf(test.Expect)
	receivedev := reflect.New(thetype)
	err := json.Unmarshal(actualBody, receivedev.Interface())
	assert.NoError(t, err)
	return receivedev.Interface()
}

// Ajax struct object
type Ajax struct {
	Method  string
	Path    string
	Body    io.Reader
	Headers map[string]string
}

// NewTestRequest is the shit
func NewTestRequest(t *testing.T, ajax Ajax) (*http.Response, []byte) {
	req, err := http.NewRequest(ajax.Method, ajax.Path, ajax.Body)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}
	for k, v := range ajax.Headers {
		req.Header[k] = []string{v}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}
	defer resp.Body.Close()

	return resp, respBody
}

// TestRequest is the shit
func TestRequest(t *testing.T, method, path string, body io.Reader) (*http.Response, []byte) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, []byte{}
	}
	defer resp.Body.Close()

	return resp, respBody
}
