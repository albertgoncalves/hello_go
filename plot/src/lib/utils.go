package utils

import (
    "fmt"
    "gonum.org/v1/plot"
)

func InitPlot() *plot.Plot {
    p, err := plot.New()
    PanicIf(err)

    return p
}

func TypeOf(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func PanicIf(err error) {
    if err != nil {
        panic(err)
    }
}
