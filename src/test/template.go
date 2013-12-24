package main

import "html/template"
import "os"
import "fmt"


func main() {
    tmpl, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
    if err != nil {
        fmt.Println(err)
    }
    err = tmpl.ExecuteTemplate(os.Stdout, "T", "replace_word")
    fmt.Println()
    if err != nil {
        fmt.Println(err)
    }
    t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.Username}}! Main Page:[{{.MainPage}}]{{end}}`)
    argMap := map[string] string { "Username": "yuki070", "MainPage": "http://weibo.com/ljq070" }
    _ = t.ExecuteTemplate(os.Stdout, "T", argMap)
    fmt.Println()

    upload, _ := template.ParseFiles("./upload.html")
    upload.Execute(os.Stdout, "Upload>>")
    fmt.Println()
}
