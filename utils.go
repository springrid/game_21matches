package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func mapDivision(array1, array2 map[int]float64) map[int]float64 {
	newArray := make(map[int]float64)
	for key := range array1 {
		newArray[key] = array1[key] / array2[key]
	}
	return newArray
}

func plotResult(array []int) {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "21 Matches winner"
	p.X.Label.Text = "Game"
	p.Y.Label.Text = "Winner"

	pts := make(plotter.XYs, len(array))

	for i, m := range array {
		pts[i].X = float64(i)
		pts[i].Y = float64(m)
	}

	// Make a line plotter and set its style.
	l, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	if err := p.Save(16*vg.Inch, 16*vg.Inch, "result.png"); err != nil {
		panic(err)
	}
}

// func sortKeysInMap(array map[int]float64) []int {
// 	var keys []int
// 	for k := range array {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)
// 	return keys
// }
