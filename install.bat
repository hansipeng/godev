@echo off
cd /d %~dp0
call WeChatAlarmServer.exe -service=install
pause