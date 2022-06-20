package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

func CMD(comand string, args ...string) (string, error) {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command(comand, args...)
	cmd.Dir = Path
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Erro ao executar comando")
	}
	cmdOut, errStd := stdOut.String(), stdErr.String()
	if errStd != "" {
		return "", errors.New(errStd)
	}
	return cmdOut, nil
}

func Organizer(name, ext string) error {
	var folderRaiz string
	for key, item := range kind {
		for _, value := range item {
			if ext == value {
				folderRaiz = key
			}
		}
	}

	// cria a pasta raiz, onde será armazenados os arquivos desse tipo
	_, err := CMD("mkdir", "-p", folderRaiz)
	if err != nil {
		return err
	}

	data := time.Now()
	ano := data.Format("2006")
	mes := data.Format("01")
	diaHora := data.Format("02-15:04:05")

	_, err = CMD("mkdir", "-p", fmt.Sprintf("%s/%s", folderRaiz, ano))
	if err != nil {
		return err
	}

	_, err = CMD("mkdir", "-p", fmt.Sprintf("%s/%s/%s", folderRaiz, ano, mes))
	if err != nil {
		return err
	}

	nameFolder := fmt.Sprintf("%s-%s", diaHora, name)

	// pasta com dia e hora e nome do arquivo
	_, err = CMD("mkdir", "-p", fmt.Sprintf("%s/%s/%s/%s", folderRaiz, ano, mes, nameFolder))
	if err != nil {
		return err
	}

	// movendo o arquivo para a nova pasta
	_, err = CMD("mv", fmt.Sprintf("%s.%s", name, ext), fmt.Sprintf("%s/%s/%s/%s", folderRaiz, ano, mes, nameFolder))
	if err != nil {
		return err
	}

	return nil
}
