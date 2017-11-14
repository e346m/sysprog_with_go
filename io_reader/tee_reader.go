package main
import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
)

func main() {
    var buffer bytes.Buffer
    reader := bytes.NewBufferString("Example of io.TeeReader\n")
    TeeReader := io.TeeReader(reader, &buffer)
    _, _ = ioutil.ReadAll(TeeReader)
    fmt.Println(buffer.String())
}
