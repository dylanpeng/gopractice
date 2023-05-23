package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	// 示例使用
	points := []Point{
		{X: 2, Y: 3},
		{X: 5, Y: 6},
		{X: 8, Y: 9},
	}

	centerPoint := Point{X: 10, Y: 10}

	newPositions := movePoints(points, centerPoint)

	for _, newPosition := range newPositions {
		fmt.Printf("New Position: (%f, %f)\n", newPosition.X, newPosition.Y)
	}
}

func movePoints(points []Point, center Point) []Point {
	var newPositions []Point

	centroid := calculateCentroid(points)
	vectors := calculateVectors(centroid, points)

	for _, vector := range vectors {
		newPosition := Point{
			X: center.X + vector.X,
			Y: center.Y + vector.Y,
		}
		newPositions = append(newPositions, newPosition)
	}

	return newPositions
}

func calculateCentroid(points []Point) Point {
	var centroid Point
	for _, point := range points {
		centroid.X += point.X
		centroid.Y += point.Y
	}
	centroid.X /= float64(len(points))
	centroid.Y /= float64(len(points))
	return centroid
}

func calculateVectors(centroid Point, points []Point) []Point {
	var vectors []Point
	for _, point := range points {
		vector := Point{
			X: centroid.X - point.X,
			Y: centroid.Y - point.Y,
		}
		vectors = append(vectors, vector)
	}
	return vectors
}
