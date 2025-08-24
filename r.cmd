@echo off
del /q output\win32\*.*

pushd cmd\win32api-gen
    go build -v .
popd
if %errorlevel% neq 0 goto :eof

cmd\win32api-gen\win32api-gen.exe
if %errorlevel% neq 0 goto :eof

pushd output\win32
    :: mod the constant, not the type
    sed -E -i "s/^\tHTTP_VERSION\b/\tHTTP_VERSION_/;" Networking.HttpServer.go
    :: mode the constant
    :: sed -E -i "s/\tRESTRICTIONS\b/\tRESTRICTIONS__/;" NetworkManagement.NetworkPolicyServer.go
    :: mod the struct field, not the function Data()
    sed -E -i "s/Data \[2\]uint64/Data_ [2]uint64/;" Networking.WebSocket.go
    :: mod the struct field, not the function Data()
    sed -E -i "s/Data \[5\]uint64/Data_ [5]uint64/;" Security.Authentication.Identity.go
    :: convert uintptr(-1) => uintptr(0xffff)
    sed -E -i "s/uintptr\(-(.)\)/uintptr(0xffff-\1+1)/;" UI.Controls.go
    :: mod the constant, not the type
    sed -E -i "s/\tVK_F\b/\tVK_F_/;" UI.Input.KeyboardAndMouse.go
popd

xcopy /y /i /q output\win32\*.* ..\go-win32api\win32\
if %errorlevel% neq 0 goto :eof

pushd ..\go-win32api\examples\propertystore
    go build -v .
popd
if %errorlevel% neq 0 goto :eof

pushd ..\go-win32api\examples\console_io
    go build -v .
popd
if %errorlevel% neq 0 goto :eof

pushd ..\go-win32api\examples\window_helloworld
    go build -v .
popd
if %errorlevel% neq 0 goto :eof
