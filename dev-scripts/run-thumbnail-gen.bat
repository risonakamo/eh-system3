@REM build and run thumbnail gen

cd %~dp0\..

go build -o output/gen_thumbnails.exe bin/gen_thumbnails/gen_thumbnails.go
output\gen_thumbnails.exe -c config2

cd %~dp0