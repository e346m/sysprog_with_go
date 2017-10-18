package main
import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer
    buffer.Write([]byte("bytes.BUffer example\n"))
    fmt.Println(buffer.String())
}
