package main

import (
    "fmt"
    "os"
)

func main() {
    userFile := "./f_w.txt"
    fout, err := os.Create(userFile)
    defer fout.Close()
    if err != nil {
        fmt.Println(userFile, err)
        return
    }

    for i := 0; i < 10; i++ {
        fout.WriteString("Hello world!\n")
        fout.Write([]byte("abcd!\n"))
    }
}
