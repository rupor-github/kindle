@echo OFF
title win64

set GOROOT=d:\go
set GOPATH=%CD%
set _TOOLS=%~dp0

set CC=gcc.exe
set CXX=g++.exe

set BITS=64
set INSTROOT=%_TOOLS%mingw%BITS%
set PKG_CONFIG_LIBDIR=%INSTROOT%\lib%BITS%\pkgconfig

set CGO_ENABLED=1
set GOHOSTOS=windows
set GOHOSTARCH=amd64
set GOOS=windows
set GOARCH=amd64

set TERM=msys

set PATH=%GOPATH%\bin;%PATH%;%GOROOT%\bin;%INSTROOT%\bin;%_TOOLS%Git\bin;%_TOOLS%SlikSvn\bin;%_TOOLS%Mercurial;%_TOOLS%bin;%_TOOLS%Bazaar
