package main

import (
	"syscall/js"

	"github.com/justinclift/webgl"
)

var gl js.Value

func main() {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "mycanvas")
	width := canvasEl.Get("clientWidth").Int()
	height := canvasEl.Get("clientHeight").Int()
	canvasEl.Call("setAttribute", "width", width)
	canvasEl.Call("setAttribute", "height", height)
	gl = canvasEl.Call("getContext", "webgl")
	if gl == js.Undefined() {
		gl = canvasEl.Call("getContext", "experimental-webgl")
	}
	// once again
	if gl == js.Undefined() {
		js.Global().Call("alert", "browser might not support webgl")
		return
	}

	// * VERTEX BUFFER *
	var verticesNative = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
	var vertices = webgl.SliceToTypedArray(verticesNative)

	// Create buffer
	vertexBuffer := gl.Call("createBuffer", webgl.ARRAY_BUFFER)

	// Bind to buffer
	gl.Call("bindBuffer", webgl.ARRAY_BUFFER, vertexBuffer)

	// Pass data to buffer
	gl.Call("bufferData", webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)

	// Unbind buffer
	gl.Call("bindBuffer", webgl.ARRAY_BUFFER, nil)

	// * INDEX BUFFER *
	var indicesNative = []uint32{
		2, 1, 0,
	}
	var indices = webgl.SliceToTypedArray(indicesNative)

	// Create buffer
	indexBuffer := gl.Call("createBuffer", webgl.ELEMENT_ARRAY_BUFFER)

	// Bind to buffer
	gl.Call("bindBuffer", webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)

	// Pass data to buffer
	gl.Call("bufferData", webgl.ELEMENT_ARRAY_BUFFER, indices, webgl.STATIC_DRAW)

	// Unbind buffer
	gl.Call("bindBuffer", webgl.ELEMENT_ARRAY_BUFFER, nil)

	// * Shaders *

	// Vertex shader source code
	vertCode := `
	attribute vec3 coordinates;
	
	void main(void) {
		gl_Position = vec4(coordinates, 1.0);
	}`

	// Create a vertex shader object
	vertShader := gl.Call("createShader", webgl.VERTEX_SHADER)

	// Attach vertex shader source code
	gl.Call("shaderSource", vertShader, vertCode)

	// Compile the vertex shader
	gl.Call("compileShader", vertShader)

	// Fragment shader source code
	fragCode := `
	void main(void) {
		gl_FragColor = vec4(0.0, 0.0, 1.0, 1.0);
	}`

	// Create fragment shader object
	fragShader := gl.Call("createShader", webgl.FRAGMENT_SHADER)

	// Attach fragment shader source code
	gl.Call("shaderSource", fragShader, fragCode)

	// Compile the fragment shader
	gl.Call("compileShader", fragShader)

	// Create a shader program object to store the combined shader program
	shaderProgram := gl.Call("createProgram")

	// Attach a vertex shader
	gl.Call("attachShader", shaderProgram, vertShader)

	// Attach a fragment shader
	gl.Call("attachShader", shaderProgram, fragShader)

	// Link both the programs
	gl.Call("linkProgram", shaderProgram)

	// Use the combined shader program object
	gl.Call("useProgram", shaderProgram)

	// * Associating shaders to buffer objects *

	// Bind vertex buffer object
	gl.Call("bindBuffer", webgl.ARRAY_BUFFER, vertexBuffer)

	// Bind index buffer object
	gl.Call("bindBuffer", webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)

	// Get the attribute location
	coord := gl.Call("getAttribLocation", shaderProgram, "coordinates")

	// Point an attribute to the currently bound VBO
	gl.Call("vertexAttribPointer", coord, 3, webgl.FLOAT, false, 0, 0)

	// Enable the attribute
	gl.Call("enableVertexAttribArray", coord)

	// * Drawing the triangle *

	// Clear the canvas
	gl.Call("clearColor", 0.5, 0.5, 0.5, 0.9)
	gl.Call("clear", webgl.COLOR_BUFFER_BIT)

	// Enable the depth test
	gl.Call("enable", webgl.DEPTH_TEST)

	// Set the view port
	gl.Call("viewport", 0, 0, width, height)

	// Draw the triangle
	gl.Call("drawElements", webgl.TRIANGLES, len(indicesNative), webgl.UNSIGNED_SHORT, 0)
}
