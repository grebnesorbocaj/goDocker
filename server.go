package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    host := os.Getenv("HOST")
    targetUrl := "http://" + host + ":3000"
    fmt.Printf("Send reqeust to nodeDocker container at route -> %s", targetUrl)
    resp, err := http.Get(targetUrl)
    if err != nil {
	fmt.Fprintf(w, "error! error! error!: %s", err)
        return 
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, "Hi there, I love %s and %s", r.URL.Path[1:], string(body))
    return
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
