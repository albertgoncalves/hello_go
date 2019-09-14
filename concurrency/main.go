package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

func deliver(response chan<- string, payload *string, n int) {
    time.Sleep(time.Duration(n) * time.Second)
    var m int
    if n < 2 {
        m = 1
    } else {
        m = n - 1
    }
    response <- fmt.Sprintf("%s%s", strings.Repeat("\t", m), *payload)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    response := make(chan string)
    go func() {
        for {
            scanner.Scan()
            payload := scanner.Text()
            go deliver(response, &payload, 2)
            go deliver(response, &payload, 4)
        }
    }()
    for {
        fmt.Println(<-response)
    }
}
