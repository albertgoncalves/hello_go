package main

import (
    F "fmt"
    I "io/ioutil"
    R "regexp"
    S "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func readFile() string {
    text, err := I.ReadFile("text")
    check(err)
    return string(text)
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
    reg, err := R.Compile("[^a-z0-9 ]+")
    check(err)
    return hist(S.Fields(reg.ReplaceAllString(S.ToLower(s), "")))
}

func main() {
    m := pipeline(readFile())
    for k, v := range m {
        F.Printf("{%s: %d}\n", k, v)
    }
}
