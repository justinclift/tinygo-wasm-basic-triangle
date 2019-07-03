This is just simple research code, directly copied from bobcob7,
to see if TinyGo wasm can be used for WebGL.

Running demo:

&nbsp; &nbsp; https://justinclift.github.io/tinygo-wasm-basic-triangle/

Original source:

&nbsp; &nbsp; https://github.com/bobcob7/wasm-basic-triangle

To compile the WebAssembly file:

    $ tinygo build -target wasm -no-debug -o docs/wasm.wasm main.go

To strip the custom name section from the end (reducing file size
further):

    $ wasm2wat docs/wasm.wasm -o docs/wasm.wat
    $ wat2wasm docs/wasm.wat -o docs/wasm.wasm
    $ rm -f docs/wasm.wat

