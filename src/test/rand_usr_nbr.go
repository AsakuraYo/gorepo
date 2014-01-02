package main

import "fmt"
import "math/rand"

func main() {
    const SIZE = 99999999
    var mark [SIZE]byte
    var third = [4]int { 3, 4, 5, 8 }

    r := rand.New(rand.NewSource(99))

    for i, j := 0, 0; i < 12000000; i, j = i + 1, (j + 1) % 4 {
        var num int
        for {
            tmp := r.Intn(SIZE)
            if mark[tmp] == 0 {
                mark[tmp] = 1
                num = tmp
                break
            }
        }
        fmt.Printf("13%d%08d\n", third[j], num)
    }
}
