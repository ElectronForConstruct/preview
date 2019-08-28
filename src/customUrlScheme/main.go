package customUrlScheme

import (
	"fmt"
	"runtime"
)

// TODO https://support.shotgunsoftware.com/hc/en-us/articles/219031308-Launching-applications-using-custom-browser-protocols#linux

func Register() bool {
	switch currentOS := runtime.GOOS; currentOS {
	case "darwin":
		fmt.Println("Support for MacOS is coming soon.")
	case "linux":
		fmt.Println("Support for Linux is coming soon.")
	case "windows":
		return WindowsRegister()
	default:
		fmt.Printf("%s.\n", currentOS)
	}
	return true
}
