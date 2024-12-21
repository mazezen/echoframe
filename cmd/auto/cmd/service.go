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

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runServiceCode()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	handlerCmd.Flags().String("service", "se", "generate service code")
}

var sf = "/internal/app/service"

func runServiceCode() {
	serviceTem := rc()

	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strService := itools.CompactStr(rootDir, sf, strings.ToLower(model.Name), "_", "service.go")
		if _, err := os.Stat(strService); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}
	checkModel(ms)

	if err := generateServiceCode(ms, string(serviceTem)); err != nil {
		log.Fatal(err)
	}
}

func rc() []byte {
	serviceTem, err := os.ReadFile("./templates/service.template")
	if err != nil {
		log.Fatal(err)
	}
	return serviceTem
}

func generateServiceCode(models []dt.ModelInfo, serviceTemplate string) error {
	serviceTmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return err
	}

	sd := itools.CompactStr(rootDir, sf)
	for _, model := range models {
		_, err = os.Stat(sd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(sd, 0777); err != nil {
				log.Fatal(err)
			}
		}

		serviceFileName := fmt.Sprintf("%s/%s_service.go", sd, strings.ToLower(model.Name))
		_ = do.GenerateFile(serviceTmpl, serviceFileName, model)
	}

	return nil
}
