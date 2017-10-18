package main
import (
    "os"
    "fmt"
    "time"
)

func main () {
    fmt.Fprint(os.Stdout, "write with os.Stdout at %v", time.Now())
}
