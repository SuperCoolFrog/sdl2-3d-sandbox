[string]$sourceDirectory = "assets"
[string]$destinationDirectory = "bin"
Copy-item -Force -Recurse $sourceDirectory -Destination $destinationDirectory
go build -o bin/simple.exe sdl2-3d-sandbox/cmd/main ; start .\bin\simple.exe 
