package sslx

import "testing"
//import "bytes"
//import "fmt"

func TestEncrypt(t *testing.T) {
    const keyStr = "gmcc1234"
    const text = "hello world liangx"
    var encryptor Encryptor
    encryptor = new(AES)
    encryptor.Encrypt([]byte(text), keyStr)
}

func TestDecrypt(t *testing.T) {
    const keyStr = "gmcc1234"
    const text = "hello world liangx"
    var encryptor Encryptor
    encryptor = &AES{}
    encryptBytes := encryptor.Encrypt([]byte(text), keyStr)
    var decryptor Decryptor
    decryptor = &AES{}
    decryptor.Decrypt(encryptBytes, keyStr)
}

func BenchmarkEncrypt(b *testing.B) {
    const keyStr    = "1234567890123456"
    const text      = "1234567890123456"
    var encryptor Encryptor
    encryptor = new(AES)
    for i := 0; i < 1000000; i++ {
        encryptor.Encrypt([]byte(text), keyStr)
    }
}

func BenchmarkDecrypt(b *testing.B) {
    const keyStr    = "1234567890123456"
    const text      = "1234567890123456"
    var decryptor Decryptor
    decryptor = new(AES)
    for i := 0; i < 1000000; i++ {
        decryptor.Decrypt([]byte(text), keyStr)
    }
}
