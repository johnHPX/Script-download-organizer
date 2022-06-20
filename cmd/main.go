package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/johnHPX/sdo/util"
)

func main() {
	for {
		files, err := ioutil.ReadDir(util.Path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			s := strings.Split(file.Name(), ".")
			if len(s) > 1 {
				err := util.Organizer(s[0], s[1])
				if err != nil {
					log.Fatal(err)
				}
				log.Println("organized file!")
			}
		}
	}

}
