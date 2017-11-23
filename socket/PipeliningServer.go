package main
import (
    "bufio"
    "time"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "io"
    "io/ioutil"
    "strings"
)

func writeToConn(sessionResponse chan chan *http.Response, conn net.Conn) {
    defer conn.Close()
    for sessionResponse := range sessionResponse {
        response := <- sessionResponse
        response.Write(conn)
        close(sessionResponse)
    }
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(dump))
    content := "Hello World\n"
    response := &http.Response {
        StatusCode: 200,
        ProtoMajor: 1,
        ProtoMinor: 1,
        ContentLength: int64(len(content)),
        Body: ioutil.NopCloser(strings.NewReader(content)),
    }
    resultReceiver <- response
}

func processSession(conn net.Conn) {
    fmt.Printf("Accept %v\n", conn.RemoteAddr())
    sessionResponses := make(chan chan *http.Response, 50)
    defer close(sessionResponses)
    go writeToConn(sessionResponses, conn)
    reader := bufio.NewReader(conn)
    for {
        conn.SetReadDeadline(time.Now().Add(5 * time.Second))
        request, err := http.ReadRequest(reader)
        if err != nil {
            neterr,ok := err.(net.Error)
            if ok && neterr.Timeout() {
                fmt.Println("Timeout")
                break
            } else if err == io.EOF {
                break
            }
            panic(err)
        }
        sessionResponse := make(chan *http.Response)
        sessionResponses <- sessionResponse
        go handleRequest(request, sessionResponse)
    }
}

func main() {
    listener, err := net.Listen("tcp", "localhost:8888")
    if err != nil {
        panic(err)
    }
    fmt.Println("server is runnging at locahost:8888")
    if err != nil {
        panic(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }
        go processSession(conn)
    }
}
