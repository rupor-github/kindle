@echo OFF
title win32

set GOROOT=d:\go
set GOPATH=%CD%
set _TOOLS=%~dp0
set _MSYS=d:\msys64\

set CC=gcc.exe
set CXX=g++.exe

set BITS=32
set INSTROOT=%_MSYS%mingw%BITS%
set PKG_CONFIG_LIBDIR=%INSTROOT%\lib\pkgconfig

set CGO_ENABLED=1
set GOHOSTOS=windows
set GOHOSTARCH=amd64
set GOOS=windows
set GOARCH=386

set TERM=msys

set PATH=%GOPATH%\bin;%PATH%;%GOROOT%\bin;%INSTROOT%\bin;%_MSYS%\usr\bin
