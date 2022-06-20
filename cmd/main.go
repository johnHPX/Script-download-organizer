package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/johnHPX/sdo/util"
)

func main() {
	for {
		// lendo os arquivos do diretório Downloads
		files, err := ioutil.ReadDir(util.Path)
		if err != nil {
			log.Fatal(err)
		}

		// percorrendo os resultados
		for _, file := range files {
			// verificado se o item encontrado é realmente um arquivo baixado, e não por exemplo uma pasta
			s := strings.Split(file.Name(), ".")
			if len(s) == 2 {
				err := util.Organizer(s[0], s[1])
				if err != nil {
					log.Fatal(err)
				}
				log.Println("organized file!")
			}
		}
	}
}
