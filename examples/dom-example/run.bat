@echo off
echo Checking Go version...
go version

echo Building WebAssembly module...
set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go

echo Copying wasm_exec.js...
for /f "tokens=*" %%i in ('go env GOROOT') do set GOROOT=%%i
echo Using GOROOT: %GOROOT%
copy "%GOROOT%\misc\wasm\wasm_exec.js" . /Y

echo Verifying wasm_exec.js...
if not exist wasm_exec.js (
    echo Error: wasm_exec.js was not copied successfully
    exit /b 1
)

echo Building Docker image...
docker build -t go-wasm-dom-example .

echo Running Docker container...
docker run -d -p 8080:80 --name go-wasm-dom go-wasm-dom-example

echo Done! Open http://localhost:8080 in your browser
echo Press any key to stop the container...
pause > nul

echo Stopping container...
docker stop go-wasm-dom
docker rm go-wasm-dom 