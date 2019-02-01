package main

import (
    "encoding/csv"
    "log"
    "os"
)

var data = [][]string{
    {"a", "b", "b"},
    {"1", "2", "3"},
}

func main() {
    const fn string = "data.csv"

    file, err := os.Create(fn)
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    w.WriteAll(data) // method includes 'defer writer.Flush()'
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
