package main

import (
    "fmt"
    "flag"
    "os"
    "net"
    "io"
)

func main () {
    var listenFlag bool
    flag.BoolVar(&listenFlag, "l", false, "listen for an incoming connection rather than connect to it")
    flag.Parse()

    if flag.NArg() != 2 {
        fmt.Println("Usage : ", os.Args[0], "[-l] hostname port")
        flag.PrintDefaults()
        os.Exit(1)
    }
    hostname := flag.Arg(0)
    port := flag.Arg(1)
    addr := net.JoinHostPort(hostname, port)

    if listenFlag == true {
        // as a server
        listener, err := net.Listen("tcp", addr)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        defer conn.Close()

        for {
            n, err := io.Copy(os.Stdout, conn)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
                break
            }
            if n == 0 {
                break
            }
        }
    } else {
        // as a client
        conn, err := net.Dial("tcp", addr)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        defer conn.Close()

        for {
            n, err := io.Copy(conn, os.Stdin)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
                break
            }
            if n == 0 {
                break;
            }
        }
    }
}
