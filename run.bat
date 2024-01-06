@REM builds and runs server

go build -o output/eh_system.exe bin/server/server.go
output\eh_system.exe -c config1