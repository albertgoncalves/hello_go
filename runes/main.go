package main

import (
    "fmt"
)

func report(element, char rune) {
    fmt.Printf("%d\t%#U    \t%t\n", element, element, element == char)
}

func main() {
    str := "\t\n aA€�"
    chs := []rune(str)
    report(chs[0], '\t')
    report(chs[1], '\n')
    report(chs[2], ' ')
    report(chs[3], 'a')
    report(chs[4], 'A')
    report(chs[5], '€')
    report(chs[6], '�')
}
