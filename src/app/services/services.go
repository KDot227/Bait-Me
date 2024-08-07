package services

import (
	"bait-me/handlers/errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var processes_path = "C:\\Program Files (x86)\\Bait-Me\\processes"

func StartServices() {
	// Check if the services are running
	// If not, start them
	var process_names = []string{"ida64.exe", "vmwareuser.exe", "vboxservice.exe", "wireshark.exe"}
	for _, process_name := range process_names {
		if !isProcessRunning(process_name) {
			initalizeServices(process_name)
			startup_service(process_name)
		}
	}
}

func StopServices() {
	// Stop the services
	var process_names = []string{"ida64.exe", "vmwareuser.exe", "vboxservice.exe", "wireshark.exe"}
	for _, process_name := range process_names {
		service_no_ext := strings.TrimSuffix(process_name, filepath.Ext(process_name))
		cmd := exec.Command("sc.exe", "stop", service_no_ext)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RemoveServices() {
	var process_names = []string{"ida64.exe", "vmwareuser.exe", "vboxservice.exe", "wireshark.exe"}
	for _, process_name := range process_names {
		service_no_ext := strings.TrimSuffix(process_name, filepath.Ext(process_name))
		cmd := exec.Command("sc.exe", "delete", service_no_ext)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func isProcessRunning(process_name string) bool {
	// Check if the process is running
	// Return true if it is, false otherwise
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.Contains(string(output), process_name) {
		return true
	}
	return false
}

func initalizeServices(process_name string) {
	files, err := os.ReadDir(processes_path)
	if err != nil {
		errors.MakeError(err.Error())
	}
	for _, file := range files {
		if file.Name() == process_name {
			cmd := exec.Command(filepath.Join(processes_path, process_name), "--mode=install")
			err := cmd.Start()
			if err != nil {
				fmt.Println(err)
			}
			process_name_no_ext := strings.TrimSuffix(process_name, filepath.Ext(process_name))
			cmd = exec.Command("sc.exe", "start", process_name_no_ext)
			cmd.Run()
		}
	}
}

func startup_service(process_name string) {
	service_no_ext := strings.TrimSuffix(process_name, filepath.Ext(process_name))
	cmd := exec.Command("sc.exe", "start", service_no_ext)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
