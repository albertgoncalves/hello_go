package main

import (
    "fmt"
)

func typeOf(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func matrix(n int, m int) [][]int {
    mat := make([][]int, n)
    row := make([]int, n*m)

    for i := 0; i < n; i++ {
        mat[i] = row[i*m : (i+1)*m]
    }

    return mat
}

func transpose(mat [][]int) [][]int {
    n := len(mat)
    m := len(mat[0])
    tr := matrix(m, n)

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            tr[j][i] = mat[i][j]
        }
    }

    return tr
}

func printAll(xs ...interface{}) {
    for _, x := range xs {
        fmt.Println(x)
    }
}

func main() {
    const n int = 5
    const m int = 3

    x := matrix(n, m)

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            x[i][j] = i
        }
    }

    printAll(typeOf(x), x, transpose(x))
}
