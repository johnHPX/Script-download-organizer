package util

import (
	"fmt"
	"log"
	"strings"
)

// extensões
var kind = map[string][]string{
	"imagens":               []string{"jpeg", "jpg", "gif", "png", "bmp", "psd", "tiff", "svg", "raw", "webp"},
	"documentos":            []string{"txt", "doc", "docx", "ppt", "pps", "xls", "xlsx"},
	"arquivos-comapactados": []string{"zip", "rar", "tar", "dep"},
	"codigos":               []string{"go", "html", "css", "js", "php", "py", "c", "jar"},
	"musicas":               []string{"mp3", "aac", "flac", "alac", "aiff", "ape", "dsd", "mqa"},
	"videos":                []string{"mkv", "mp4", "wmv", "avi", "flv", "ogg", "avchd", "mpg"},
}

var Path string

func init() {
	user, err := CMD("whoami")
	if err != nil {
		log.Fatal(err)
	}
	u := strings.TrimSpace(user)
	Path = fmt.Sprintf("/home/%s/Downloads", u)
}
