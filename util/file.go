package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

// CMD é responsável pela execução dos comandos Bash do linux. ela retorna dois valores, a saída e um error.
func CMD(comand string, args ...string) (string, error) {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command(comand, args...)
	// defini aonde o comando Bash será executado, no nosso caso na pasta "Downloads".
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

// Organizer responsável por organizar os arquivos em pastas, nomeadas em categorias, ano, mês e pasta própria dos arquivos.
func Organizer(name, ext string) error {
	// verificando se a extenção do arquivo está presente no Kind
	var folderRaiz string
	var exists bool
	for key, item := range kind {
		for _, value := range item {
			if ext == value {
				folderRaiz = key
				exists = true
			}
		}
	}
	// caso não exista, o arquivo vai ser colocado na pasta "outros".
	if !exists{
		folderRaiz = "outros"
	}

	// criar a pasta raiz, onde vai ficar o grupo na qual o arquivo faz parte.
	_, err := CMD("mkdir", "-p", folderRaiz)
	if err != nil {
		return err
	}

	// definindo variavies de tempo
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

	// nome da pasta pessoal do arquivo
	nameFolder := fmt.Sprintf("%s-%s", diaHora, name)

	// pasta com dia, hora e nome do arquivo
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
