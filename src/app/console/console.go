package console

import (
	"syscall"
	"unsafe"
)

const SW_HIDE = 0
const SW_SHOW = 1

var (
	modUser32      = syscall.NewLazyDLL("user32.dll")
	procFindWindow = modUser32.NewProc("FindWindowW")
	procShowWindow = modUser32.NewProc("ShowWindow")
)

func HideConsole() {
	hwnd, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ConsoleWindowClass"))), 0)
	if hwnd != 0 {
		procShowWindow.Call(hwnd, SW_HIDE)
	}
}

func ShowConsole() {
	hwnd, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ConsoleWindowClass"))), 0)
	if hwnd != 0 {
		procShowWindow.Call(hwnd, SW_SHOW)
	}
}
