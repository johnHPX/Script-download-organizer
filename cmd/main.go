package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/johnHPX/sdo/util"
)

func main() {
	log.Println("==========================================================")
	control := "yes"
	for control == "yes" {
		// lendo os arquivos do diretório Downloads
		files, err := os.ReadDir(util.Path)
		if err != nil {
			log.Fatal(err)
		}
		var cont int = 0
		// percorrendo os resultados
		for _, file := range files {
			var f string
			if util.OStype == "win" {
				f = fmt.Sprintf("%s\\%s", util.Path, file.Name())
			} else {
				f = fmt.Sprintf("%s/%s", util.Path, file.Name())
			}
			// verificado se o item encontrado é realmente um arquivo, e não uma pasta
			info, err := os.Stat(f)
			if err != nil {
				log.Fatal(err)
			}

			if info.IsDir() {
				fmt.Println("It is a directory")
				continue
			} else {
				nameFileDivDot := strings.Split(file.Name(), ".")
				extFile := nameFileDivDot[len(nameFileDivDot)-1]
				nameFileWithOutExt := strings.Join(nameFileDivDot[0:len(nameFileDivDot)-1], ".")
				err := util.Organizer(nameFileWithOutExt, extFile)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Arquivo organizado!")
				cont += 1
			}
		}
		log.Println("Total de arquivos organizados:", cont)
		log.Println("Quer organizar novamente a pasta Downloads? [yes/no]")
		fmt.Scan(&control)
	}
	log.Println("==========================================================")
}
