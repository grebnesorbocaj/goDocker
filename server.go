package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Println("Send reqeust to nodeDocker container at route -> http://nodeDocker:3000")
    resp, err := http.Get("http://nodedocker:3000/")
    if err != nil {
        fmt.Println("error! error! error!")
        fmt.Println(err)
        return 
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("Received: " + string(body))
    return
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
