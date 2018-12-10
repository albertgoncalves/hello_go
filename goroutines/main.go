package main

import (
    "fmt"
    "time"
)

func say(message string, pause time.Duration) {
    for {
        fmt.Println(message)
        time.Sleep(pause)
    }
}

func main() {
    go say("Marco!", 2*time.Second)
    say("Polo!", 3*time.Second)
}
