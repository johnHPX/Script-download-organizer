package util

import (
	"fmt"
	"log"
	"strings"
)

// Kind é um map com as chaves, representando os tipos de arquivos e os slices de string com os valores desses tipos
var kind = map[string][]string{
	"imagens":               []string{"jpeg", "jpg", "gif", "png", "bmp", "psd", "tiff", "svg", "raw", "webp"},
	"documentos":            []string{"txt", "doc", "docx", "ppt", "pps", "xls", "xlsx"},
	"arquivos-comapactados": []string{"zip", "rar", "tar", "dep"},
	"codigos":               []string{"go", "html", "css", "js", "php", "py", "c", "jar"},
	"musicas":               []string{"mp3", "aac", "flac", "alac", "aiff", "ape", "dsd", "mqa"},
	"videos":                []string{"mkv", "mp4", "wmv", "avi", "flv", "ogg", "avchd", "mpg"},
}

// Path é a variavel que define o Caminho até a pasta Downloads
var Path string

// init Configura o Path, para cada usuario de um sistema linux.
func init() {
	user, err := CMD("whoami")
	if err != nil {
		log.Fatal(err)
	}
	u := strings.TrimSpace(user)
	Path = fmt.Sprintf("/home/%s/Downloads", u)
}
