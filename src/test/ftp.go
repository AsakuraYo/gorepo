package main

import (
    ftp "code.google.com/p/ftp4go"
)

func main() {
    conn := ftp.NewFTP(0)
    conn.Connect("127.0.0.1", 21, "")
}
