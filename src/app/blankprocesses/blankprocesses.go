package blankprocesses

import "os"

var paths = []string{
	"C:\\Windows\\vboxmrxnp.dll",
	"C:\\Windows\\System32\\vmGuestLib.dll",
}

func MakeBlankProcesses() {
	for _, path := range paths {
		data := make([]byte, 0)
		_, err := os.Stat(path)
		if err == nil {
			continue
		}

		file, err := os.Create(path)
		if err != nil {
			continue
		}
		defer file.Close()

		file.Write(data)
	}
}
