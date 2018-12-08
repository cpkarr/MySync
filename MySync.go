package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var path string
	var localErr error

	path, localErr = exec.LookPath("rsync")

	if localErr != nil {
		fmt.Println("Cannot find rsync")
		os.Exit(404)
	}

	localCmd := exec.Command(path, "-a", "/Users/ckarr/go/src", "/Users/ckarr/Destination/")
	localErr = localCmd.Run()
	if localErr == nil {
		fmt.Println("rsync called successfully")
	} else {
		fmt.Println("rsync called unuccessfully")
	}
}
