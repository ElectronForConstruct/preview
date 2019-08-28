package main

import (
	"../customUrlScheme"
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Return the current path
func runningPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func main() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "efc://") {
		url := strings.Replace(os.Args[1], "efc://", "", 1)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the path to your project root: ")
		directory, _ := reader.ReadString('\n')
		directory = strings.TrimSuffix(directory, "\n")
		directory = strings.TrimSuffix(directory, "\r")
		println(directory)

		startPreview(url, directory)
	} else {

		previewCommand := flag.NewFlagSet("preview", flag.ExitOnError)
		registerCommand := flag.NewFlagSet("register", flag.ExitOnError)
		unregisterCommand := flag.NewFlagSet("unregister", flag.ExitOnError)

		urlPreviewCommand := previewCommand.String("url", "", "The URL you want to preview (Required)")

		if len(os.Args) < 2 {
			fmt.Println("You must specify a subcommand")
			os.Exit(1)
		}

		switch os.Args[1] {
		case "preview":
			previewCommand.Parse(os.Args[2:])
		case "register":
			registerCommand.Parse(os.Args[2:])
		case "unregister":
			unregisterCommand.Parse(os.Args[2:])
		default:
			flag.PrintDefaults()
			os.Exit(1)
		}

		if previewCommand.Parsed() {
			// Required Flags
			if *urlPreviewCommand == "" {
				previewCommand.PrintDefaults()
				os.Exit(1)
			}

			startPreview(*urlPreviewCommand, runningPath())
		}

		if registerCommand.Parsed() {
			customUrlScheme.Register()
		}

		if unregisterCommand.Parsed() {
			// work
			fmt.Println("unregisterCommand")
		}
	}
}

func startPreview(url string, currentDir string) {

	reader := bufio.NewReader(os.Stdin)

	cmd := exec.Command("efc", "preview", `--url="`+url+`"`)
	cmd.Dir = currentDir

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Println("Node.js and @efc/cli module must be installed before using this binary!")
		}
		reader.ReadString('\n')
		os.Exit(0)
	}

	// there is an error:

	cmd.Stdout.Write([]byte("Press enter to exit...\n"))
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
