package privileges

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func EnsureAdmin() {
	if !isAdmin() {
		err := runAsAdmin()
		if err != nil {
			fmt.Println("Error while trying to get administrative privileges:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func isAdmin() bool {
	// Check for administrative privileges
	var sid *windows.SID
	sid, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}

	token := windows.Token(0)
	admin, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return admin
}

func runAsAdmin() error {
	// Get the path to the executable
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	verb := "runas"
	// Relaunch the program with elevated privileges
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process -Verb %s -FilePath '%s'", verb, exe))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
