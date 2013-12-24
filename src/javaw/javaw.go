package main

import "fmt"
import "os"
import "io/ioutil"

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "<encryptFile>")
        os.Exit(1)
    }
    orgPath := os.Args[1]
    javaPath := orgPath + ".java"
    os.Rename(orgPath, javaPath)
    content, err := ioutil.ReadFile(javaPath)
    if err != nil {
        fmt.Println(javaPath, ": ReadFile ERROR.", err.Error())
        os.Exit(1)
    }
    txtPath := orgPath + ".txt"
    if err = ioutil.WriteFile(txtPath, content, 0664); err != nil {
        fmt.Println(javaPath, ": WriteFile ERROR.", err.Error())
        os.Exit(1)
    }
    os.Remove(javaPath)
}

