package main

import (
	"bytes"
	"log"
	"os/exec"
)

func fileDetection(path string) bool {

	cmd := exec.Command("fuser", path)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Printf("error listing processes which is opening file : %v", err)
		return false
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Printf("out: %s err: %s", outStr, errStr)
	return true
}

func main() {
	// check if target file is opened by other process
	// if the file do not exists, it will return it as not in use
	path := "/home/matrix/.xinputrc"
	if fileInUse := fileDetection(path); fileInUse {
		log.Printf("file is in use")
	} else {
		log.Printf("file is not in use, can be deleted safely")
	}
}
