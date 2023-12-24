package main

import (
	"errors"
	"fmt"
	"image/color"
	"main/pkg"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

type XYs struct {
	X []float64
	Y []float64
}

func plotHarmonic(h pkg.HarCon) (XYs, error) {
	p := plot.New()
	p.Title.Text = h.Name

	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Height"

	phase := degreesToRadians(h.PhaseGMT)
	amplitude := h.Amplitude
	speed := degreesToRadians(h.Speed)
	if amplitude <= 0 {
		return XYs{}, errors.New("empty constituent") // skip plotting this one
	}
	//cycle := int(math.Round((200 * math.Pow(2*math.Pi, 2))) / speed)
	cycle := 400000
	var pts plotter.XYs = make(plotter.XYs, cycle)
	var XYs = XYs{}
	for i := 0; i < cycle; i++ {
		// figure out what x should be
		x := float64(i) / (100 * 2 * math.Pi)
		y := amplitude * math.Sin(speed*float64(x)+phase)
		XYs.X = append(XYs.X, x)
		XYs.Y = append(XYs.Y, y)
		pts[i].X = x
		pts[i].Y = y
	}

	// Add the points to the plot.
	l, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)
	p.X.Min = 0
	p.X.Max = float64(cycle) / (100 * 2 * math.Pi)
	p.Y.Min = -15
	p.Y.Max = 13

	// Save the plot to a PNG file.

	if err := p.Save(19*vg.Inch, 10*vg.Inch, h.Name+"_points.png"); err != nil {
		panic(err)
	}
	fmt.Println("Saved " + h.Name + "_points.png")
	return XYs, nil
}

func main() {
	station := "9452210"
	HarmonicConstituents, err := pkg.GetHarmonicConstituent(station)
	if err != nil {
		panic(err)
	}
	var sumXYs plotter.XYs = make(plotter.XYs, 400000)
	for i, HarCon := range HarmonicConstituents.HarCons {
		xys, err := plotHarmonic(HarCon)
		if err != nil {
			continue
		}
		for i := range xys.X {
			sumXYs[i].X = xys.X[i]
			sumXYs[i].Y += xys.Y[i]
		}

		// Add the points to the plot.
		p := plot.New()
		p.Title.Text = "Sum of all Harmonics"
		p.X.Label.Text = "Time?"
		p.Y.Label.Text = "Height"
		l, err := plotter.NewLine(plotter.XYs(sumXYs))
		if err != nil {
			panic(err)
		}
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Color = color.RGBA{B: 255, A: 255}
		p.Add(l)
		p.X.Min = 0
		p.X.Max = float64(400000) / (100 * 2 * math.Pi)
		p.Y.Min = -15
		p.Y.Max = 13

		// Save the plot to a PNG file.
		if err := p.Save(19*vg.Inch, 10*vg.Inch, fmt.Sprint(i)+HarCon.Name+"_sum_points.png"); err != nil {
			panic(err)
		}
	}
}
