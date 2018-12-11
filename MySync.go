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

// First make sure that the rsync tool is installed. Should never be a problem, but better safe than sorry
	path, localErr = exec.LookPath("rsync")

	if localErr != nil {
		fmt.Println("Cannot find rsync")
		os.Exit(404)
	}
// Next, start doing the sync/backups

//Here is the first backup
	localCmd := exec.Command(path, "-a", "/Volumes/4TB_RAID1/BackupToCloud", "/Volumes/Backup/rsyncMacBackups")
	_, localErr = localCmd.StderrPipe()
	if localErr != nil {
		log.Fatal(localErr)
		os.Exit(-1)
	}
	localErr = localCmd.Run()
	if localErr != nil {
		log.Fatal(localErr)
		os.Exit(-1)
	}

	fmt.Println("first backup ran successfully")

//Here is the second backup

	localCmd = exec.Command(path, "-a", "/Volumes/4TB_RAID1/PhotosLibraries", "/Volumes/Backup/rsyncMacBackups")
	_, localErr = localCmd.StderrPipe()
	if localErr != nil {
		log.Fatal(localErr)
		os.Exit(-1)
	}
	localErr = localCmd.Run()
	if localErr != nil {
		log.Fatal(localErr)
		os.Exit(-1)
	}
	fmt.Println("second backup ran successfully")
}
