package main

import (
    "fmt"
)

type Obj struct {
    num int
}

func (o *Obj) Num() int {
    return o.num
}

func (o *Obj)SetNum(num int) {
    o.num = num
}

func say (obj *Obj) {
    fmt.Println(obj.Num())
    obj.SetNum(obj.Num() + 1)
}

func main () {
    obj := &Obj{num:1}
    say(obj)
    say(obj)
    say(obj)
}
