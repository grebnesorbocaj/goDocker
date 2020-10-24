package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("Send reqeust to nodeDocker container at route -> http://nodeDocker:3000")
    resp, err := http.Get("http://nodedocker:3000/")
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
