@echo OFF

set _TOOLS=%~dp0
set _MSYS=d:\msys64\
set INSTROOT=%_TOOLS%armkpw
set GOHOSTOS=linux
set GOHOSTARCH=arm

if "x%CROSS_SAVE_PATH%"=="x" set CROSS_SAVE_PATH=%PATH%

if x%1 == x goto fail
if x%2 == x goto fail

if x%1 == xwindows goto win
if x%1 == xlinux   goto lin
goto fail 

:lin

if x%2 == xarm goto linarm
goto fail 

:win

if x%2 == xamd64 goto win64
if x%2 == x386    goto win32
goto fail 

:win64

title arm Cross for %1 %2
set INSTROOT_FOR_TARGET=%_MSYS%mingw64

set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1

set CC=arm-mingw_kpw-linux-gnueabi-gcc.exe
set CXX=arm-mingw_kpw-linux-gnueabi-g++.exe
set PKG_CONFIG_LIBDIR=%INSTROOT%\lib\pkgconfig
set PATH=%INSTROOT%\bin;%CROSS_SAVE_PATH%

set CC_FOR_TARGET=gcc.exe
set CXX_FOR_TARGET=g++.exe
set PKG_CONFIG_LIBDIR_FOR_TARGET=%INSTROOT_FOR_TARGET%\lib64\pkgconfig
set PATH_FOR_TARGET=%INSTROOT_FOR_TARGET%\bin;%CROSS_SAVE_PATH%

echo. 
echo Setting Cross for %1 %2

goto fin

:win32

title arm Cross for %1 %2
set INSTROOT_FOR_TARGET=%_MSYS%mingw32

set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

set CC=arm-mingw_kpw-linux-gnueabi-gcc.exe
set CXX=arm-mingw_kpw-linux-gnueabi-g++.exe
set PKG_CONFIG_LIBDIR=%INSTROOT%\lib\pkgconfig
set PATH=%INSTROOT%\bin;%CROSS_SAVE_PATH%

set CC_FOR_TARGET=gcc.exe
set CXX_FOR_TARGET=g++.exe
set PKG_CONFIG_LIBDIR_FOR_TARGET=%INSTROOT_FOR_TARGET%\lib\pkgconfig
set PATH_FOR_TARGET=%INSTROOT_FOR_TARGET%\bin;%CROSS_SAVE_PATH%

echo. 
echo Setting Cross for %1 %2

goto fin

:linarm

title arm Cross for %1 %2
set INSTROOT_FOR_TARGET=%_TOOLS%armkpw

set GOOS=linux
set GOARCH=arm
set CGO_ENABLED=1

set CC=arm-mingw_kpw2-linux-gnueabi-gcc.exe
set CXX=arm-mingw_kpw2-linux-gnueabi-g++.exe
set PKG_CONFIG_LIBDIR=%INSTROOT%\lib\pkgconfig
set PATH=%INSTROOT%\bin;%CROSS_SAVE_PATH%

set CC_FOR_TARGET=arm-mingw_kpw2-linux-gnueabi-gcc.exe
set CXX_FOR_TARGET=arm-mingw_kpw2-linux-gnueabi-g++.exe
set PKG_CONFIG_LIBDIR_FOR_TARGET=
set PATH_FOR_TARGET=%INSTROOT_FOR_TARGET%\bin;%CROSS_SAVE_PATH%

echo. 
echo Setting Cross for %1 %2

goto fin

:fail
echo "buildgo_cross GOOS GOARCH!"

:fin
echo.