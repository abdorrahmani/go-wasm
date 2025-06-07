@echo off
echo Building WebAssembly module...
set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go

echo Copying wasm_exec.js...
for /f "tokens=*" %%i in ('go env GOROOT') do set GOROOT=%%i
copy "%GOROOT%\misc\wasm\wasm_exec.js" .

echo Building Docker image...
docker build -t go-wasm-example .

echo Running Docker container...
docker run -d -p 8080:80 --name go-wasm go-wasm-example

echo Done! Open http://localhost:8080 in your browser
echo Press any key to stop the container...
pause > nul

echo Stopping container...
docker stop go-wasm
docker rm go-wasm 