package main

import (
	"fmt"
	"log"
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

	localCmd := exec.Command(path, "-a", "/Volumes/4TB_RAID1/BackupToCloud", "/Volumes/Backup/rsyncMacBackups")
	_, err := localCmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	localErr = localCmd.Run()
	if localErr == nil {
		fmt.Println("first backup ran successfully")
		localCmd := exec.Command(path, "-a", "/Volumes/4TB_RAID1/PhotosLibraries", "/Volumes/Backup/rsyncMacBackups")
		_, err := localCmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		localErr = localCmd.Run()
		if localErr == nil {
			fmt.Println("second backup ran successfully")
		}
	}
}
