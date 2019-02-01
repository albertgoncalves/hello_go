package main

import (
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    U "lib"
    "math/rand"
)

// $ go get gonum.org/v1/plot
// https://github.com/gonum/plot/wiki/Example-plots

func quartilePlot(i float64, x plotter.Values) *plotter.QuartPlot {
    q, err := plotter.NewQuartPlot(i, x)
    U.PanicIf(err)
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

    p := U.InitPlot()

    p.Title.Text = "Quartile plots"
    p.Y.Label.Text = "Values"

    p.Add(
        quartilePlot(0, uniform),
        quartilePlot(1, normal),
        quartilePlot(2, expon),
    )

    p.NominalX(
        "Uniform\nDistribution",
        "Normal\nDistribution",
        "Exponential\nDistribution",
    )

    U.PanicIf(p.Save(6*vg.Inch, 8*vg.Inch, "pngs/quartile.png"))
}
