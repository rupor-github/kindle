set _TOOLS=%~dp0

d:
cd d:\go\src

ren .\make.bat .\make.orig
copy %_TOOLS%make.cmd .

call %_TOOLS%buildgo_cross_amd64.cmd windows amd64
call .\make.cmd
call %_TOOLS%buildgo_cross_amd64.cmd windows 386
call .\make.cmd --no-clean
call %_TOOLS%buildgo_cross_amd64.cmd linux arm
call .\make.cmd --no-clean
