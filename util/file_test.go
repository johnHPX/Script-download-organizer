package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestHelloName(t *testing.T) {
	var args_new []string
	args_new = append(args_new, "mv")
	args_new = append(args_new, "'.\\Firefox Installer.exe'")
	args_new = append(args_new, ".\\outros\\2023\\10\\Firefox-Installer\\")
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("powershell", args_new...)
	// defini aonde o comando Bash será executado, no nosso caso na pasta "Downloads".
	cmd.Dir = "C:\\Users\\jonatas freitas\\Downloads"
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}
