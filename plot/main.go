package main

import (
    "fmt"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "math/rand"
)

// $ go get gonum.org/v1/plot
// https://github.com/gonum/plot/wiki/Example-plots

// examine types of objects
// fmt.Println(typeOf(uniform) ...
func typeOf(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func alarm(err error) {
    if err != nil {
        panic(err)
    }
}

func initPlot() *plot.Plot {
    p, err := plot.New()
    alarm(err)
    return p
}

func quartilePlot(i float64, x plotter.Values) *plotter.QuartPlot {
    q, err := plotter.NewQuartPlot(i, x)
    alarm(err)
    return q
}

func main() {
    rand.Seed(int64(0))
    const n = 10000

    uniform := make(plotter.Values, n)
    normal := make(plotter.Values, n)
    expon := make(plotter.Values, n)

    for i := 0; i < n; i++ {
        uniform[i] = rand.Float64()
        normal[i] = rand.NormFloat64()
        expon[i] = rand.ExpFloat64()
    }

    p := initPlot()

    p.Title.Text = "Quartile plots"
    p.Y.Label.Text = "Values"

    p.Add(quartilePlot(0, uniform),
        quartilePlot(1, normal),
        quartilePlot(2, expon))

    p.NominalX("Uniform\nDistribution",
        "Normal\nDistribution",
        "Exponential\nDistribution")

    alarm(p.Save(6*vg.Inch, 8*vg.Inch, "pngs/quartile.png"))
}
