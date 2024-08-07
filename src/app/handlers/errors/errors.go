package errors

import (
	"golang.org/x/sys/windows"
)

func MakeError(text string) int32 {
	// Make a box with the text
	out_int, err := windows.MessageBox(0, windows.StringToUTF16Ptr(text), windows.StringToUTF16Ptr("Error"), windows.MB_OK)
	if err != nil {
		panic(err)
	}
	return out_int
}
