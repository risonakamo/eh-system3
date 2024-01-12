@REM builds and runs server

cd %~dp0\..

go build -o output/eh_system.exe bin/server/server.go
output\eh_system.exe -c config1

cd %~dp0