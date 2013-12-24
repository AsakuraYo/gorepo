package main

/*
#include <openssl/aes.h>

void encrypt(const char *key) {
    AES_KEY aes;
    AES_set_encrypt_key(key, 16 * 8, &aes);
}
*/
import (
    "C"
    "fmt"
)

func main() {
	str := C.CString("Hello world")
    fmt.Println(str)
    C.encrypt(str)
}
