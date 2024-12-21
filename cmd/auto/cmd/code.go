/*
Copyright © 2024 NAME mazezen
*/
package cmd

import (
	"fmt"
	"github.com/mazezen/echoframe/cmd/auto/do"
	"github.com/mazezen/echoframe/cmd/auto/dt"
	"github.com/mazezen/echoframe/utils"
	"github.com/mazezen/itools"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
	"strings"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCode()
	},
}

var rootDir string

func init() {
	rootCmd.AddCommand(codeCmd)
	rootDir, _ = utils.FindProjectRoot()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	codeCmd.Flags().String("code", "c", "generate handler, service, dao code")
}

func runCode() {
	handlerTem, serviceTem, dtoTem := rh(), rc(), rd()
	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strHandler := itools.CompactStr(rootDir, hf, strings.ToLower(model.Name), "_", "handler.go")
		if _, err := os.Stat(strHandler); os.IsNotExist(err) {
			ms = append(ms, model)
		}
		strService := itools.CompactStr(rootDir, sf, strings.ToLower(model.Name), "_", "service.go")
		if _, err := os.Stat(strService); os.IsNotExist(err) {
			ms = append(ms, model)
		}
		strDto := itools.CompactStr(rootDir, df, strings.ToLower(model.Name), "_", "dto.go")
		if _, err := os.Stat(strDto); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}
	checkModel(models)
	if err := generateCode(ms, string(handlerTem), string(serviceTem), string(dtoTem)); err != nil {
		log.Fatal(err)
	}
}

// generateHandlerCode generates handler code based on the model information
func generateCode(models []dt.ModelInfo, handlerTemplate, serviceTemplate, dtoTemplate string) error {
	handlerTmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	serviceTmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return err
	}

	dtoTmpl, err := template.New("dao").Parse(dtoTemplate)
	if err != nil {
		return err
	}

	hd := itools.CompactStr(rootDir, hf)
	sd := itools.CompactStr(rootDir, sf)
	dd := itools.CompactStr(rootDir, df)
	var (
		handlerFileName string
		serviceFileName string
		dtoFileName     string
	)
	for _, model := range models {
		// Generate handler code
		_, err = os.Stat(hd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(hd, 0777); err != nil {
				log.Fatal(err)
			}
		}

		handlerFileName = fmt.Sprintf("%s/%s_handler.go", hd, strings.ToLower(model.Name))
		_ = do.GenerateFile(handlerTmpl, handlerFileName, model)

		// Generate service code
		_, err = os.Stat(sd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(sd, 0777); err != nil {
				panic(err)
			}
		}
		serviceFileName = fmt.Sprintf("%s/%s_service.go", sd, strings.ToLower(model.Name))
		_ = do.GenerateFile(serviceTmpl, serviceFileName, model)

		// Generate dto code
		_, err = os.Stat(dd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(dd, 0777); err != nil {
				panic(err)
			}
		}
		dtoFileName = fmt.Sprintf("%s/%s_dto.go", dd, strings.ToLower(model.Name))
		_ = do.GenerateFile(dtoTmpl, dtoFileName, model)
	}
	return nil
}

func checkModel(models []dt.ModelInfo) {
	if len(models) == 0 {
		log.Fatal("no models")
	}
}
