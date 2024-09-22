// simulation/simulation.go
package simulation

import (
    "runtime"
    "sync"
    "time"
    "go-graphics/entities"
)

// initializes simulation, possibly using multiple threads.
func InitSimulation(wg *sync.WaitGroup, points []entities.Point) {
    runtime.GOMAXPROCS(runtime.NumCPU()) // use all available CPU cores

    // runs simulation in a separate goroutine
    go func() {
        defer wg.Done()
        simulateWorld(points)
    }()
}

// simulateWorld updates the position of each point based on its velocity
func simulateWorld(points []entities.Point) {
    for {
        for i := range points {
            points[i].X += points[i].VelX
            points[i].Y += points[i].VelY
            points[i].Z += points[i].VelZ

            // boundary checks to keep points within a certain range
            if points[i].X < -1.0 || points[i].X > 1.0 {
                points[i].VelX *= -1 // reflect velocity off x boundary
            }
            if points[i].Y < -1.0 || points[i].Y > 1.0 {
                points[i].VelY *= -1 // reflect velocity off y boundary
            }
            if points[i].Z < -1.0 || points[i].Z > 1.0 {
                points[i].VelZ *= -1 // reflect velocity off z boundary
            }
        }
        // roughly 60 fps
        time.Sleep(time.Millisecond * 16)
    }
}
