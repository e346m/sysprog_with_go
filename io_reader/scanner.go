package main
import (
    "bufio"
    "fmt"
    "strings"
)

var source = `Line 1
Line 2
Line 3
`

func main() {
    scanner := bufio.NewScanner(strings.NewReader(source))
    for scanner.Scan() {
        fmt.Printf("%#v\n", scanner.Text())
    }
}
