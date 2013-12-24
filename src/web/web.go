package main

import (
    "net/http"
    "fmt"
    "log"
    "html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    // fmt.Println("Form params:", r.Form)
    // fmt.Println("path =", r.URL.Path)
    // fmt.Println("scheme =", r.URL.Scheme)
    // for k, v := range r.Form {
    //     fmt.Println(k, "=>", v)
    // }
    // fmt.Println("=================")
    // fmt.Fprintf(w, "Hello liangx")
    fmt.Println(r.Method, r.URL.Path, r.Form)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./upload.html")
        t.Execute(w, nil)
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/", 200)
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/upload", upload)
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatal("ListenAndServe", err)
    }
}
