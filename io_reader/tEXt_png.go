package main

import (
    "bytes"
    "hash/crc32"
    "encoding/binary"
    "fmt"
    "io"
    "os"
)

func readChunks(file *os.File) []io.Reader {
    var chunks []io.Reader

    file.Seek(8, 0)
    var offset int64 = 8
    for {
        var length int32
        err := binary.Read(file, binary.BigEndian, &length)
        if err == io.EOF {
            break
        }
        chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
        offset, _ = file.Seek(int64(length+8), 1)
    }
    return chunks
}

func textChunk(text string) io.Reader {
    byteData := []byte(text)
    var buffer bytes.Buffer
    binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
    buffer.WriteString("tEXt")
    buffer.Write(byteData)
    crc := crc32.NewIEEE()
    io.WriteString(crc, "tEXt")
    binary.Write(&buffer, binary.BigEndian, crc.Sum32())
    return &buffer
}

func main() {
    file, err := os.Open("Lenna.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    newFile, err := os.Create("Lenna2.png")
    if err != nil {
        panic(err)
    }
    defer newFile.Close()
    chunks := readChunks(file)
    io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
    io.Copy(newFile, chunks[0])
    io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
    for _, chunk := range chunks[1:] {
        io.Copy(newFile, chunk)
    }
}
