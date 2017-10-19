package main

import (
    "os"
    "io"
    "strings"
    "archive/zip"
)

func main() {
    file, err := os.Create("sample.zip")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    zipWriter := zip.NewWriter(file)
    defer zipWriter.Close()

    a, err := zipWriter.Create("a.txt")
    if err != nil {
        panic(err)
    }
    io.Copy(a, strings.NewReader("text file No.1"))

    b, err := zipWriter.Create("b.txt")
    if err != nil {
        panic(err)
    }
    io.Copy(b, strings.NewReader("text file No.2"))
}
