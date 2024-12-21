/*
Copyright © 2024 NAME mazezen
*/
package cmd

import (
	"fmt"
	"github.com/mazezen/echoframe/cmd/auto/do"
	"github.com/mazezen/echoframe/cmd/auto/dt"
	"github.com/mazezen/itools"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
	"strings"
)

// handlerCmd represents the handler command
var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runHandlerCode()
	},
}

func init() {
	rootCmd.AddCommand(handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	handlerCmd.Flags().String("handler", "ha", "generate handler code")
}

var hf = "/internal/app/handler"

func runHandlerCode() {
	handlerTem := rh()

	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strHandler := itools.CompactStr(rootDir, hf, strings.ToLower(model.Name), "_", "handler.go")
		if _, err := os.Stat(strHandler); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}

	checkModel(ms)
	if err := generateHandlerCode(ms, string(handlerTem)); err != nil {
		log.Fatal(err)
	}
}

func rh() []byte {
	handlerTem, err := os.ReadFile("./templates/handler.template")
	if err != nil {
		log.Fatal(err)
	}
	return handlerTem
}

func generateHandlerCode(models []dt.ModelInfo, handlerTemplate string) error {
	handlerTmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}
	hd := itools.CompactStr(rootDir, hf)
	for _, model := range models {
		_, err = os.Stat(hd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(hd, 0777); err != nil {
				log.Fatal(err)
			}
		}

		handlerFileName := fmt.Sprintf("%s/%s_handler.go", hd, strings.ToLower(model.Name))
		_ = do.GenerateFile(handlerTmpl, handlerFileName, model)
	}
	return nil
}
