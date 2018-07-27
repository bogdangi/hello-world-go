package main

import (
    "os"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_HelloWorld(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    HelloWorld(res, req)

    name, err := os.Hostname()
    exp := fmt.Sprintf("Hello World from %s", name)
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
