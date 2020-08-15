package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

type logWriter struct {}

func main() {
    res, err := http.Get("http://www.google.com")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    lw := logWriter{}

    io.Copy(lw, res.Body)
}

// Custom Writer function implementing the Writer interface & built in Write function
func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Wrote Bytes:", len(bs))
    return len(bs), nil
}
