@echo off
echo Setting up WebAssembly test environment...

REM Build and run the test container
echo Building test container...
docker build -t go-wasm-tests -f Dockerfile.test ..

echo Running tests...
docker run -d -p 8080:8080 --name wasm-tests go-wasm-tests

echo.
echo Test environment is running!
echo Open http://localhost:8080/test.html in your browser to view test results
echo.
echo Press any key to stop the test environment...
pause > nul

echo Cleaning up...
docker stop wasm-tests
docker rm wasm-tests 