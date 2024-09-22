// entities/point.go
package entities

import (
    "math/rand"
)

// point in 3D space with position, color, and velocity (likely need to do quaternions or something)
type Point struct {
	X, Y, Z float64 // position
	R, G, B float64 // color
	VelX, VelY, VelZ float64 // velocity
}

// generates a list of random points with random velocities (to be replaced)
func GenerateRandomPoints(n int, min, max float64) []Point {
	points := make([]Point, n)
	for i := 0; i < n; i++ {
			points[i] = Point{
					X: rand.Float64()*(max-min) + min,
					Y: rand.Float64()*(max-min) + min,
					Z: rand.Float64()*(max-min) + min,
					R: rand.Float64(),
					G: rand.Float64(),
					B: rand.Float64(),
					VelX: rand.Float64()*0.02 - 0.01,
					VelY: rand.Float64()*0.02 - 0.01,
					VelZ: rand.Float64()*0.02 - 0.01,
			}
	}
	return points
}

