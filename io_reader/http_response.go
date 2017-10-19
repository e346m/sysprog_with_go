package main
import(
    "bufio"
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
)
func main() {
    conn, err := net.Dial("tcp", "example.org:80")
    if err != nil {
        panic(err)
    }
    conn.Write([]byte("Get /HTTP/1.0\r\nHost: example.org\r\n\r\n"))
    res, err := http.ReadResponse(bufio.NewReader(conn), nil)
    fmt.Println(res.Header)
    defer res.Body.Close()
    io.Copy(os.Stdout, res.Body)
}
