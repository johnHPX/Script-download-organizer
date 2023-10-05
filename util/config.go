package util

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Kind é um map com as chaves, representando os tipos de arquivos e os slices de string com os valores desses tipos
var kind = map[string][]string{
	"imagens":              {"jpeg", "jpg", "gif", "png", "bmp", "psd", "tiff", "svg", "raw", "webp"},
	"documentos":           {"txt", "doc", "docx", "ppt", "pps", "xls", "xlsx"},
	"arquivos-compactados": {"zip", "rar", "tar", "dep"},
	"codigos":              {"go", "html", "css", "js", "php", "py", "c", "jar"},
	"musicas":              {"mp3", "aac", "flac", "alac", "aiff", "ape", "dsd", "mqa"},
	"videos":               {"mkv", "mp4", "wmv", "avi", "flv", "ogg", "avchd", "mpg"},
}

// Path é a variavel que define o Caminho até a pasta Downloads
var Path string

// OSType é a variavel que contém o nome do sistema operacional
var OStype string

// init Configura o Path, para cada usuario de um sistema linux.
func init() {

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Seu sistema operacional é Windows!")
		var arg []string
		arg = append(arg, "whoami")
		user, err := CMD(arg)
		if err != nil {
			log.Fatal(err)
		}
		user_slice := strings.TrimSpace(user)
		u := strings.Split(user_slice, "\\")[1]
		Path = fmt.Sprintf("C:\\Users\\%s\\Downloads", u)
		OStype = "win"
	case "darwin":
		fmt.Println("Seu sistema operacional é MAC OS!")
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Seu sistema operacional é da familia Linux!")
		var arg []string
		arg = append(arg, "whoami")
		user, err := CMD(arg)
		if err != nil {
			log.Fatal(err)
		}
		u := strings.TrimSpace(user)
		Path = fmt.Sprintf("/home/%s/Downloads", u)
		OStype = "linux"
	}

}
