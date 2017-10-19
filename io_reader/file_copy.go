package main
import (
    "io"
    "flag"
    "os"
)

func main() {
    source := flag.String("s", "old.txt", "source file")
    destination := flag.String("d", "new.txt", "destination file")
    flag.Parse()

    s_file, err := os.Open(*source)
    if err != nil {
        panic(err)
    }
    defer s_file.Close()

    w_file, err := os.Create(*destination)
    if err != nil {
        panic(err)
    }
    io.Copy(w_file, s_file)
}
