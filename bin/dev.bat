@ECHO OFF

rem Starts the app. It doesn't reload on Windows

cd %~dp0\..

set solitaire_encryption_key=TEMP_SECRET_KEY
set GOEXPERIMENT=jsonv2

echo "Windows doesn't allow reloading... sorry"
@ECHO ON
go.exe build -gcflags "all=-N -l" -o build/debug/solitaire.exe .
build\debug\solitaire.exe -v --addr=0.0.0.0 all
