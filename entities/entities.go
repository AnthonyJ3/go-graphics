// entities/entities.go
package entities

import "fmt"

var points []Point

// initializes entities in the scene and returns the points
func InitEntities() []Point {
    fmt.Println("Initializing entities...")
    points = GenerateRandomPoints(1000, -1.0, 1.0)
    fmt.Printf("Generated %d random points\n", len(points))
    return points
}

// returns the list of points
func GetPoints() []Point {
    return points
}

