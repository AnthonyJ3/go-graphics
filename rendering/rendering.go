// rendering/rendering.go
package rendering

import (
    "github.com/go-gl/gl/v4.6-core/gl"
    "github.com/go-gl/glfw/v3.3/glfw"
    "go-graphics/entities"
    "log"
    "unsafe"
)

// initializes the GLFW and OpenGL contexts
func Init() *glfw.Window {
    if err := glfw.Init(); err != nil {
        log.Fatalln("failed to initialize GLFW:", err)
    }
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OpenGL 4.x
    glfw.WindowHint(glfw.ContextVersionMinor, 6) // OpenGL 4.6
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.Resizable, glfw.False)

    window, err := glfw.CreateWindow(800, 600, "3D Graphics", nil, nil)
    if err != nil {
        log.Fatalln("failed to create window:", err)
    }
    window.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        log.Fatalln("failed to initialize OpenGL:", err)
    }

    width, height := window.GetFramebufferSize()
    gl.Viewport(0, 0, int32(width), int32(height))

    // close callback to make sure the close event works
    window.SetCloseCallback(func(w *glfw.Window) {
        w.SetShouldClose(true)
    })

    return window
}

// clears the buffer for the next frame
func Clear() {
	gl.ClearColor(1.0, 1.0, 1.0, 1.0) 
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// swaps the double buffer to display the rendered frame
func SwapBuffers(window *glfw.Window) {
    window.SwapBuffers()
}

// renders the given points using modern OpenGL with VBOs
func RenderPoints(points []entities.Point) {
    // flatten the points array into a single float32 array
    var vertices []float32
    for _, point := range points {
        vertices = append(vertices, float32(point.X), float32(point.Y), float32(point.Z))
    }

    // generate and bind a Vertex Array Object (VAO)
    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

    // generate and bind a Vertex Buffer Object (VBO)
    var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)

    // define layout of the vertex data (currently 3 floats per vertex)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
    gl.EnableVertexAttribArray(0)

    // draw points
    gl.DrawArrays(gl.POINTS, 0, int32(len(points)))

    // cleanup
    gl.DisableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindVertexArray(0)

    // delete the VBO and VAO after rendering (in some cases, could keep them for multiple frames)
    gl.DeleteBuffers(1, &vbo)
    gl.DeleteVertexArrays(1, &vao)
}

