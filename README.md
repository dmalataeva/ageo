# S2 Cell ID Converter

This project is a single-page application that converts latitude/longitude pairs and an S2 cell level to an S2 cell ID. It uses WebAssembly to bridge the gap between the Go library (`github.com/golang/geo`) and the JavaScript front-end.

## Prerequisites

*   Go (version 1.18 or later)

## Compilation

1.  Navigate to `go/` and compile the Go code to WebAssembly:

    ```bash
    cd go
    GOOS=js GOARCH=wasm go build -o ../scripts/main.wasm main.go
    ```

## Running Locally

1.  Serve the files using the provided Express server:

    ```bash
    node server.js
    ```

3.  Open your browser and navigate to `http://localhost:8000` to see the application in action.

## Usage

1.  Enter the latitude, longitude, and S2 level in the input fields.
2.  Click the "Convert" button.
3.  The S2 cell ID will be displayed in the result area.

## Troubleshooting

`TypeError: Failed to execute 'compile' on 'WebAssembly': Incorrect response MIME type. Expected 'application/wasm'.`

Your server is unable to recognize the application type and use the right content header. Try updating `etc/mime.types with:
```
image/webp                  webp
```
