// +build windows

package customUrlScheme

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"path/filepath"
)

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

