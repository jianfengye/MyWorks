package main

import (
	"bytes"
	"os/exec"
	"log"
)

func Run(dir string, args ...string) ([]byte, error) {
	var buf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	
	cmd.Dir = dir
	cmd.Stdout = &buf
	cmd.Stderr = cmd.Stdout
	log.Println("run cmd:", cmd.Args)

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Wait for the command.  Clean up,
	err := cmd.Wait()
	return buf.Bytes(), err
}
