package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func TextEditor(content []byte) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return []byte{}, fmt.Errorf("please set the EDITOR environment variable")
	}
	f, err := ioutil.TempFile("", "pd_")
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	if err := f.Chmod(0600); err != nil {
		return []byte{}, err
	}
	if _, err := f.Write(content); err != nil {
		return []byte{}, err
	}
	if err := f.Close(); err != nil {
		return []byte{}, err
	}
	cmdParts := strings.Fields(editor)
	cmd := exec.Command(cmdParts[0], append(cmdParts[1:], f.Name())...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return []byte{}, err
	}
	ct, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return []byte{}, err
	}
	return ct, nil
}
