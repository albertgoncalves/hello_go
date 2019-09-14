package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    response := make(chan string)
    go func() {
        for {
            scanner.Scan()
            payload := scanner.Text()
            go func() {
                time.Sleep(3 * time.Second)
                response <- payload
            }()
        }
    }()
    for {
        fmt.Println("\t", <-response)
    }
}
