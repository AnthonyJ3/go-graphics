// main.go
package main

import (
    "fmt"
    "runtime"
    "sync"

    "go-graphics/entities"
    "go-graphics/rendering"
    "go-graphics/simulation"
    //"github.com/go-gl/gl/v4.6-core/gl"
    "github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
    runtime.LockOSThread()
}

func cleanup() {
    fmt.Println("Cleaning up...")
}

func main() {
    defer cleanup()

    // initialize rendering system (GLFW and OpenGL)
    window := rendering.Init()
    defer glfw.Terminate()

    // initialize simulation & passes points to simulation
    points := entities.InitEntities()
    var wg sync.WaitGroup
    wg.Add(1)
    simulation.InitSimulation(&wg, points)

    // rendering loop
    for !window.ShouldClose() {
        rendering.Clear()
        rendering.RenderPoints(points) 
        rendering.SwapBuffers(window)
        glfw.PollEvents()
    }
    wg.Wait()
}

