package util

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// CMD é responsável pela execução dos comandos Bash do linux. ela retorna dois valores, a saída e um error.
func CMD(args []string) (string, error) {
	var stdOut, stdErr bytes.Buffer
	var cmd *exec.Cmd
	if OStype == "win" {
		cmd = exec.Command("powershell", args...)
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}

	// Define onde o comando Bash será executado, no nosso caso na pasta "Downloads".
	cmd.Dir = Path
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	cmdOut, errStd := stdOut.String(), stdErr.String()
	if errStd != "" {
		return "", errors.New(errStd)
	}
	return cmdOut, nil
}

func organizer_win(folderRaiz, name, ext string) error {
	log.Println("Criando a pasta raiz, caso não exista...")
	// criar a pasta raiz, onde vai ficar o grupo na qual o arquivo faz parte.
	args := []string{"mkdir", "-Force", folderRaiz}
	_, err := CMD(args)
	if err != nil {
		return err
	}

	log.Println("Definindo horário...")
	// definindo variavies de tempo
	data := time.Now()
	ano := data.Format("2006")
	mes := data.Format("01")
	dia := data.Format("02")

	log.Println("Criando a pasta dia-mês-ano, caso não exista...")
	args[2] = fmt.Sprintf("%s/%s-%s-%s", folderRaiz, dia, mes, ano)
	_, err = CMD(args)
	if err != nil {
		return err
	}

	// nome da pasta pessoal do arquivo
	nameFolder := strings.Replace(name, " ", "-", -1)

	// log.Println("Criando a pasta dia e nome do arquivo, caso não exista...")
	// // pasta com dia e nome do arquivo
	// args[2] = fmt.Sprintf("%s/%s/%s/%s/%s", folderRaiz, ano, mes, dia, nameFolder)
	// _, err = CMD(args)
	// if err != nil {
	// 	return err
	// }

	// Movendo o arquivo para a nova pasta
	log.Println("Movendo o arquivo da pasta Downloads para a nova pasta...")
	args_new := []string{"mv", fmt.Sprintf("'.\\%s.%s'", name, ext), fmt.Sprintf(".\\%s\\%s\\%s\\%s\\%s\\", folderRaiz, dia, mes, ano, nameFolder)}
	_, err = CMD(args_new)
	if err != nil {
		return err
	}

	return nil
}

func organizer_linux(folderRaiz, name, ext string) error {
	log.Println("Criando a pasta raiz, caso não exista...")
	// criar a pasta raiz, onde vai ficar o grupo na qual o arquivo faz parte.
	args := []string{"mkdir", "-p", folderRaiz}
	_, err := CMD(args)
	if err != nil {
		return err
	}

	log.Println("Definindo horário...")
	// definindo variavies de tempo
	data := time.Now()
	ano := data.Format("2006")
	mes := data.Format("01")
	dia := data.Format("02")

	log.Println("Criando a pasta dia-mês-ano, caso não exista...")
	args[2] = fmt.Sprintf("%s/%s-%s-%s", folderRaiz, dia, mes, ano)
	_, err = CMD(args)
	if err != nil {
		return err
	}

	// nome da pasta pessoal do arquivo
	nameFolder := strings.Replace(name, " ", "-", -1)

	// log.Println("Criando a pasta dia e nome do arquivo, caso não exista...")
	// // pasta com dia e nome do arquivo
	// args[2] = fmt.Sprintf("%s/%s/%s/%s/%s", folderRaiz, ano, mes, dia, nameFolder)
	// _, err = CMD(args)
	// if err != nil {
	// 	return err
	// }

	// Movendo o arquivo para a nova pasta
	log.Println("Movendo o arquivo da pasta Downloads para a nova pasta...")
	args_new := []string{"mv", fmt.Sprintf("%s.%s", name, ext), fmt.Sprintf("%s/%s-%s-%s/%s", folderRaiz, dia, mes, ano, nameFolder)}
	_, err = CMD(args_new)
	if err != nil {
		return err
	}

	return nil
}

// Organizer responsável por organizar os arquivos em pastas, nomeadas em categorias, ano, mês e pasta própria dos arquivos.
func Organizer(name, ext string) error {
	log.Println("------------------------------------------------")
	log.Println("Analisando a extensão do arquivo...")
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
	if !exists {
		folderRaiz = "outros"
	}

	log.Println("Este arquivo vai ser organizado na pasta:", folderRaiz)

	if OStype == "win" {
		organizer_win(folderRaiz, name, ext)
	} else if OStype == "linux" {
		organizer_linux(folderRaiz, name, ext)
	}

	return nil
}
