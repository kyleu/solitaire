@ECHO OFF
rem Content managed by Project Forge, see [projectforge.md] for details.

rem Starts the app. It doesn't reload on Windows

cd %~dp0\..

echo "Windows doesn't allow reloading... sorry"
@ECHO ON
go.exe build -gcflags "all=-N -l" -o build/debug/solitaire.exe .
build\debug\solitaire.exe -v --addr=0.0.0.0 all
