package main

import (
    "crypto/rand"
    "os"
    "io"
)

func main() {
    file, err := os.Create("rand.txt")
    if err != nil {
        panic(err)
    }
    rand_stream := rand.Reader
    io.CopyN(file, rand_stream, 1024)
}
