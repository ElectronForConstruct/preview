package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func runningPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]

		// fmt.Println(runningPath())

		/* --- */
		npm := "npm"
		if runtime.GOOS == "windows" {
			npm = "npm.cmd"
		}

		cmd := exec.Command(npm, "run", "start", "--", url)
		cmd.Dir = runningPath()

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		// Execute the command
		if err := cmd.Run(); err != nil {
			// log.Panic(err)
			log.Println(err)
		}

		log.Println(stdBuffer.String())
		/* --- */

	} else {
		fmt.Print("You must pass an url as first argument!")
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
