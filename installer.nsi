;This is 100% AI generated if you couldn't tell
; Define the name of the installer and the output file
OutFile "Bait-Me-Installer.exe"

; Define the icon for the installer
Icon "icon.ico"

; Define the installation directory
InstallDir "$PROGRAMFILES\Bait-Me"

; Request application privileges for installation
RequestExecutionLevel admin

; Define the uninstaller
UninstallText "This will uninstall Bait-Me."
UninstallIcon "icon.ico"

; Define the sections for installation and uninstallation
Section "Install"

  ; Set the output path to the installation directory
  SetOutPath "$INSTDIR"

  ; Copy the main executable
  File "Bait-Me.exe"

  ; Create the processes directory and copy the files
  SetOutPath "$INSTDIR\processes"
  File "processes\ida64.exe"
  File "processes\vboxservice.exe"
  File "processes\vmwareuser.exe"
  File "processes\wireshark.exe"

  ; Create a shortcut in the startup folder
  CreateShortCut "$SMSTARTUP\Bait-Me.lnk" "$INSTDIR\Bait-Me.exe" "" "$INSTDIR\Bait-Me.exe"

  ; Create a shortcut in the Start Menu
  CreateDirectory "$SMPROGRAMS\Bait-Me"
  CreateShortCut "$SMPROGRAMS\Bait-Me\Bait-Me.lnk" "$INSTDIR\Bait-Me.exe" "" "$INSTDIR\Bait-Me.exe"

  ; Create a shortcut on the desktop
  CreateShortCut "$DESKTOP\Bait-Me.lnk" "$INSTDIR\Bait-Me.exe" "" "$INSTDIR\Bait-Me.exe"

  ; Write the uninstaller
  WriteUninstaller "$INSTDIR\Uninstall.exe"

  ; Add registry entries to show in "Add or Remove Programs"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "DisplayName" "Bait-Me"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "UninstallString" "$INSTDIR\Uninstall.exe"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "InstallLocation" "$INSTDIR"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "DisplayIcon" "$INSTDIR\Bait-Me.exe"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "Publisher" "KDot227"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "DisplayVersion" "1.0"
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "NoModify" 1
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me" "NoRepair" 1

SectionEnd

Section "Uninstall"

  ; Kill Bait-Me.exe if it's running
  ExecWait 'taskkill /IM "Bait-Me.exe" /F'

  ; Stop and delete services for each file in the processes directory
  SetOutPath "$INSTDIR\processes"
  ClearErrors
  FindFirst $0 $1 "$INSTDIR\processes\*.*"
  loop:
  IfErrors done
  StrCpy $2 $1
  StrCpy $3 $2 0 "."
  ExecWait "sc stop $3"
  ExecWait "sc delete $3"
  FindNext $0 $1
  Goto loop
  done:
  FindClose $0

  ; Delete the installed files
  Delete "$INSTDIR\Bait-Me.exe"
  Delete "$INSTDIR\processes\ida64.exe"
  Delete "$INSTDIR\processes\vboxservice.exe"
  Delete "$INSTDIR\processes\vmwareuser.exe"
  Delete "$INSTDIR\processes\wireshark.exe"
  
  ; Remove the processes directory
  RMDir "$INSTDIR\processes"

  ; Delete the uninstaller
  Delete "$INSTDIR\Uninstall.exe"

  ; Remove the installation directory
  RMDir "$INSTDIR"

  ; Delete the shortcut from the startup folder
  Delete "$SMSTARTUP\Bait-Me.lnk"

  ; Delete the Start Menu folder and shortcut
  Delete "$SMPROGRAMS\Bait-Me\Bait-Me.lnk"
  RMDir "$SMPROGRAMS\Bait-Me"

  ; Delete the desktop shortcut
  Delete "$DESKTOP\Bait-Me.lnk"

  ; Remove registry entries
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Bait-Me"

SectionEnd
