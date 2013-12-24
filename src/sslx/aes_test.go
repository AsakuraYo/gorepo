package sslx

import "testing"
import "bytes"
//import "fmt"

func TestEncrypt(t *testing.T) {
    const keyStr = "gmcc1234"
    const text = "hello world liangx"
    var encryptor Encryptor
    encryptor = new(AES)
    encryptor.Encrypt(bytes.NewBufferString(text).Bytes(), bytes.NewBufferString(keyStr).Bytes())
}

func TestDecrypt(t *testing.T) {
    const keyStr = "gmcc1234"
    const text = "hello world liangx"
    var encryptor Encryptor
    encryptor = &AES{}
    encryptBytes := encryptor.Encrypt(bytes.NewBufferString(text).Bytes(), bytes.NewBufferString(keyStr).Bytes())
    var decryptor Decryptor
    decryptor = &AES{}
    decryptor.Decrypt(encryptBytes, bytes.NewBufferString(keyStr).Bytes())
}
