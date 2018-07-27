package main

import (
    "os"
    "fmt"
    "net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    fmt.Fprint(res, "Hello World from ", name)
}

func main() {
    http.HandleFunc("/", HelloWorld)
    http.ListenAndServe(":3000", nil)
}
