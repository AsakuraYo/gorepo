// web project main.go
package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path, r.Form)
	var message string = ""
	if r.Method == "POST" {
		r.ParseMultipartForm(16)
		uploadfile, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer uploadfile.Close()
		fmt.Println(handler.Header)
        dir, err := os.Open("./files")
        if err != nil {
            if os.IsNotExist(err) {
                fmt.Println("./files dir is not exist")
            }
            return
        }
        defer dir.Close()
		localfile, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer localfile.Close()
		if _, err := io.Copy(localfile, uploadfile); err == nil {
			//file, err := os.Open("./files/"+handler.Filename)
			fileBytes, _ := ioutil.ReadFile("./files/" + handler.Filename)
			md5Bytes := md5.Sum(fileBytes)
			var hexBytes [md5.Size * 2]byte
			hex.Encode(hexBytes[:], md5Bytes[:])
			message = handler.Filename + " upload successfully. MD5: " + strings.ToUpper(bytes.NewBuffer(hexBytes[:]).String())
		} else {
			message = handler.Filename + " upload failed!!!"
		}

	}
	t, _ := template.ParseFiles("./index.html")
	t.Execute(w, message)
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe("127.0.0.1:9090", nil); err != nil {
		fmt.Println(err)
	}
}
