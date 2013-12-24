package main

/*
#cgo linux LDFLAGS: -lpthread
#cgo CFLAGS: -Wall
#include <desini/desini.h>

void myprint(const char *s) {
    printf("myprint: %s\n", s);
}
*/
import "C"

import "unsafe"

func main() {
	str := C.CString("Hello world")
	C.myprint(str)
	C.free(unsafe.Pointer(str))
}
