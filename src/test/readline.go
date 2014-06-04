package main

import (
//    "bufio"
//    "os"
//    "strings"
    "fmt"
    "time"
//    "crypto/aes"
    "sslx"
    "encoding/base64"
)

const (
    PASSWORD = "123456"
)

// use go routinue to encrypt and make it spend different time with `delay`
func encrypt(str string, delay time.Duration) chan string {
    result := make(chan string)
    go func() {
        time.Sleep(time.Millisecond * delay)
        cipher := &sslx.AES{}
        result <- base64.StdEncoding.EncodeToString(cipher.Encrypt([]byte(str), PASSWORD))
    }()
    return result
}

func async_encrypt(fields []string) {
    size := len(fields)
    chans := make([]chan string, size)
    for i := 0; i < size; i++ {
        chans[i] = encrypt(fields[i], time.Duration(i % 3 + 1) * 500)
    }

    var counter = 0
    for {
        for i := 0; i < size; i++ {
            select {
            case fields[i] = <-chans[i]:
                fmt.Println(i, fields[i])
                counter++
            default:
                // make it no block
            }
        }
        if counter >= size {
            break
        }
    }
}

func main() {
    //file, err := os.Open("/tmp/data")
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //defer file.Close()
    //scanner := bufio.NewScanner(file)
    //for scanner.Scan() {
    //    fields := strings.Split(scanner.Text(), "&&")
    //    for _, a := range fields {
    //        fmt.Println(a)
    //    }
    //}
    //if err := scanner.Err(); err != nil {
    //    fmt.Println(err)
    //}


    arr := []string { "1111", "2222", "3333", "4444", "5555" }
    async_encrypt(arr)
    fmt.Println(arr)

}
