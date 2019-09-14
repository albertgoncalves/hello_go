package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func deliver(response chan<- string, payload string, n int) {
    time.Sleep(time.Duration(n) * time.Second)
    response <- payload
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    a := make(chan string)
    b := make(chan string)
    go func() {
        for {
            scanner.Scan()
            payload := scanner.Text()
            go deliver(a, payload, 2)
            go deliver(b, payload, 4)
        }
    }()
    for {
        select {
        case response := <-a:
            fmt.Println("\t", response)
        case response := <-b:
            fmt.Println("\t\t", response)
        }
    }
}
