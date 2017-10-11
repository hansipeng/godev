@echo off
cd /d %~dp0
call WeChatAlarmServer.exe -service=uninstall
pause