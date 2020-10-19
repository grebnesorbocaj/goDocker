package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Println("Send reqeust to nodeDocker at localhost:8888")
    resp, err := http.Get("http://localhost:8888/")
    if err != nil {
        fmt.Println("error! error! error!")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("Received: " + string(body))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
