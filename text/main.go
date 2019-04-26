package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    S "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func readStdin() string {
    data, err := ioutil.ReadAll(os.Stdin)
    check(err)
    return string(data)
}

func hist(xs []string) map[string]int {
    m := make(map[string]int)
    for _, x := range xs {
        if val, ok := m[x]; ok {
            m[x] = val + 1
        } else {
            m[x] = 1
        }
    }
    return m
}

func pipeline(s string) map[string]int {
    reg, err := regexp.Compile("[^a-z0-9 ]+")
    check(err)
    return hist(S.Fields(reg.ReplaceAllString(S.ToLower(s), "")))
}

func handleError() {
    if err := recover(); err != nil {
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
}

func main() {
    defer handleError()
    m := pipeline(readStdin())
    for k, v := range m {
        fmt.Printf("{%s: %d}\n", k, v)
    }
}
