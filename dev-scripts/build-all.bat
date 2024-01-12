@REM builds all programs into output folder

cd %~dp0\..
go build -o output/eh_system.exe bin/server/server.go
go build -o output/gen_thumbnails.exe bin/gen_thumbnails/gen_thumbnails.go

cd %~dp0