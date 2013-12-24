// Package of openssl.
package sslx

import "crypto/aes"
//import "log"
//import "strings"

type Encryptor interface {
    Encrypt(source, key []byte) (output []byte)
}

type Decryptor interface {
    Decrypt(source, key []byte) (output []byte)
}

type AES struct {
}

func (a *AES)Encrypt(source, key []byte) (output []byte) {
    aesKeyBytes := make([]byte, aes.BlockSize)
    copy(aesKeyBytes, key)

    cipher, err := aes.NewCipher(aesKeyBytes)
    if err != nil {
        panic("aes.NewCipher: " + err.Error())
    }

    outLen := len(source)
    if len(source) % aes.BlockSize != 0 {
        outLen = (len(source) / aes.BlockSize + 1) * aes.BlockSize
    }
    sourceBytes := make([]byte, outLen)
    copy(sourceBytes, source[:])

    for i := 0; i < outLen; i = i + aes.BlockSize {
        outTemp := make([]byte, aes.BlockSize)
        cipher.Encrypt(outTemp, sourceBytes[i : i + aes.BlockSize])
        output = append(output, outTemp...)
    }
    return output
}

func (a *AES)Decrypt(source, key []byte) (output []byte) {
    aesKeyBytes := make([]byte, aes.BlockSize)
    copy(aesKeyBytes, key)

    cipher, err := aes.NewCipher(aesKeyBytes)
    if err != nil {
        panic("aes.NewCipher: " + err.Error())
    }

    inLen := len(source)
    for i := 0; i < inLen; i = i + aes.BlockSize {
        outTemp := make([]byte, aes.BlockSize)
        cipher.Decrypt(outTemp, source[i : i + aes.BlockSize])
        output = append(output, outTemp...)
    }
    return output
}

