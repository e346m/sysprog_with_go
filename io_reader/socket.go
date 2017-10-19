package main
import (
    "io"
    "os"
    "net"
)
func main() {
    conn, err := net.Dial("tcp", "example.org:80")
    if err != nil {
        panic(err)
    }
    conn.Write([]byte("GET /HTTP/1.0\r\nHost: example.org\r\n\r\n"))
    io.Copy(os.Stdout, conn)
}
