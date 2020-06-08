package main

import (
	// "crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	
	// hasher := md5.New()
	// b := []byte{}
	// io.WriteString(hasher, "test")
	// fmt.Printf("Result: %x\n", hasher.Sum(b))
	// fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher := sha1.New()
    io.WriteString(hasher, "test")
    b := []byte{}
    fmt.Printf("Result: %x\n", hasher.Sum(b))
    fmt.Printf("Result: %d\n", hasher.Sum(b))
    //
    hasher.Reset()
    data := []byte("We shall overcome!")
    n, err := hasher.Write(data)
    if n!=len(data) || err!=nil {
        log.Printf("Hash write error: %v / %v", n, err)
    }
    checksum := hasher.Sum(b)
    fmt.Printf("Result: %x\n", checksum)
}