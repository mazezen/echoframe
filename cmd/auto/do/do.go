package do

import (
	"github.com/mazezen/echoframe/cmd/auto/dt"
	"html/template"
	"log"
	"os"
)

// GenerateFile
// return 1 file create success
// return 2 file already exist
// return -1 file create error
func GenerateFile(tmpl *template.Template, fileName string, model dt.ModelInfo) int {
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		return 2
	}
	file, err := os.Create(fileName)
	if err != nil {
		return -1
	}
	defer file.Close()
	if err := tmpl.Execute(file, model); err != nil {
		log.Printf("create file err: %v\n", err)
		return -1
	}
	return 1
}
