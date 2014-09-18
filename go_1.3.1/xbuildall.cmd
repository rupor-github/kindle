set _TOOLS=%~dp0

d:
cd d:\go\src

ren .\make.bat .\make.orig
copy %_TOOLS%make.cmd .

call %_TOOLS%buildgo_cross.cmd windows amd64
call .\make.cmd
call %_TOOLS%buildgo_cross.cmd windows 386
call .\make.cmd --no-clean
call %_TOOLS%buildgo_cross.cmd linux arm
call .\make.cmd --no-clean
