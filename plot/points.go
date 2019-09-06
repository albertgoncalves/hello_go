package main

import (
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "gonum.org/v1/plot/vg/draw"
    "image/color"
    U "lib"
    "math/rand"
)

func randomPoints(n int) plotter.XYs {
    pts := make(plotter.XYs, n)

    for i := range pts {
        if i == 0 {
            pts[i].X = rand.Float64()
        } else {
            pts[i].X = pts[i-1].X + rand.Float64()
        }
        pts[i].Y = pts[i].X + 10*rand.Float64()
    }

    return pts
}

func scatter(x plotter.XYs) *plotter.Scatter {
    s, err := plotter.NewScatter(x)
    U.PanicIf(err)

    s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

    return s
}

func line(x plotter.XYs) *plotter.Line {
    l, err := plotter.NewLine(x)
    U.PanicIf(err)

    l.LineStyle.Width = vg.Points(1)
    l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
    l.LineStyle.Color = color.RGBA{B: 255, A: 255}

    return l
}

func both(x plotter.XYs) (*plotter.Line, *plotter.Scatter) {
    lpLine, lpPoints, err := plotter.NewLinePoints(x)
    U.PanicIf(err)

    lpLine.Color = color.RGBA{G: 255, A: 255}
    lpPoints.Shape = draw.PyramidGlyph{}
    lpPoints.Color = color.RGBA{R: 255, A: 255}

    return lpLine, lpPoints
}

func main() {
    rand.Seed(int64(0))
    const n = 15
    const k = 7

    p := U.InitPlot()

    p.Title.Text = "Points Example"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    p.Add(plotter.NewGrid())

    s := scatter(randomPoints(n))
    l := line(randomPoints(n))
    lpLine, lpPoints := both(randomPoints(n))

    p.Add(s, l, lpLine, lpPoints)
    p.Legend.Add("scatter", s)
    p.Legend.Add("line", l)
    p.Legend.Add("line points", lpLine, lpPoints)

    U.PanicIf(p.Save(k*vg.Inch, k*vg.Inch, "out/points.png"))
}
