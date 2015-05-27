@echo OFF
title linuxARM (kindle pw2)

set GOROOT=d:\go
set GOPATH=%CD%
set _TOOLS=%~dp0
set _MSYS=d:\msys64\

set CC=arm-mingw_kpw-linux-gnueabi-gcc.exe
set CXX=arm-mingw_kpw-linux-gnueabi-g++.exe

set INSTROOT=%_TOOLS%armkpw
set PKG_CONFIG_LIBDIR=

set CGO_ENABLED=1
set GOHOSTOS=windows
set GOHOSTARCH=386
set GOOS=linux
set GOARCH=arm
REM set GOARM=7

set TERM=msys

set PATH=%GOPATH%\bin;%PATH%;%GOROOT%\bin;%INSTROOT%\bin;%_MSYS%\usr\bin
