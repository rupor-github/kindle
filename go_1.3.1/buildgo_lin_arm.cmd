@echo OFF
title linuxARM (kindle pw2)

set GOROOT=d:\go
set GOPATH=%CD%
set _TOOLS=%~dp0

set CC=arm-mingw_kpw2-linux-gnueabi-gcc.exe
set CXX=arm-mingw_kpw2-linux-gnueabi-g++.exe

set INSTROOT=%_TOOLS%armkpw2
set PKG_CONFIG_LIBDIR=

set CGO_ENABLED=1
set GOHOSTOS=windows
set GOHOSTARCH=386
set GOOS=linux
set GOARCH=arm
set GOARM=7

set PATH=%GOPATH%\bin;%PATH%;%GOROOT%\bin;%INSTROOT%\bin;%TOOLS%SlikSvn\bin;%_TOOLS%Git\bin;%_TOOLS%Mercurial;%_TOOLS%bin;%_TOOLS%Bazaar
