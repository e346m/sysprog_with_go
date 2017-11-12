package main
import (
    "bufio"
    "fmt"
    "strings"
)

var source =`Line 1
Line 2
Line 3
`
func main() {
    reader := bufio.NewReader(strings.NewReader(source))
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Printf("%#v\n", line)
    }
}
