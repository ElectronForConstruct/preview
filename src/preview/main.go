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
	"strings"
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

		cmd := exec.Command("efc", "preview", url)
		cmd.Dir = runningPath()

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		// Execute the command
		if err := cmd.Run(); err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("Node.js and @efc/cli module must be installed before using this binary!")
			} else {
				log.Println(err)
			}
			os.Exit(0)
		}

		log.Println(stdBuffer.String())
		/* --- */

	} else {
		fmt.Print("You must pass an url as first argument!")
		os.Exit(0)
	}
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

}
