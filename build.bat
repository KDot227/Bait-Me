@echo off
setlocal enabledelayedexpansion

if "%1" == "release" (
    set "file_path=Bait-Me"
) else (
    set "file_path=bin"
)

cd /d %~dp0

if not exist %file_path% (
    mkdir %file_path%
)

pushd src\app

go mod tidy

if "%1" == "release" (
    go build -ldflags "-s -w" -o ..\..\%file_path%\Bait-Me.exe .
) else (
    go build -o ..\..\%file_path%\Bait-Me.exe .
)

popd

pushd src\mini-apps

for /d %%d in (*) do (
    pushd %%d
    if "%1" == "release" (
        go build -ldflags "-s -w" -o ..\..\..\%file_path%\processes\%%d.exe .
    ) else (
        go build -o ..\..\..\%file_path%\processes\%%d.exe .
    )
    popd
)

popd

if "%1" == "release" (
    copy icon.ico %file_path%
    copy installer.nsi %file_path%
    makensis %file_path%\installer.nsi
    del %file_path%\installer.nsi
    move %file_path%\Bait-Me-Installer.exe setup.exe
    rmdir /s /q %file_path%
)

endlocal
exit /b 0