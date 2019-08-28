package customUrlScheme

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)
import "golang.org/x/sys/windows/registry"

// TODO https://support.shotgunsoftware.com/hc/en-us/articles/219031308-Launching-applications-using-custom-browser-protocols#linux

func WindowsRegister() bool {
	classesKey, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Classes", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Open classesKey: ", err)
	}
	defer classesKey.Close()

	efcKey, _, err := registry.CreateKey(classesKey, "efc", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Create subkey: ", err)
	}

	err = efcKey.SetStringValue("URL Protocol", "")
	if err != nil {
		log.Fatal("Create classesKey: ", err)
	}

	shellKey, _, err := registry.CreateKey(efcKey, "shell", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Create subkey: ", err)
	}
	openKey, _, err := registry.CreateKey(shellKey, "open", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Create subkey: ", err)
	}
	commandKey, _, err := registry.CreateKey(openKey, "command", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Create subkey: ", err)
	}

	filename, _ := filepath.Abs(os.Args[0])

	err = commandKey.SetStringValue("", "\""+filename+"\" \"%1\"")
	if err != nil {
		log.Fatal("Create classesKey: ", err)
	}

	return true
}

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
