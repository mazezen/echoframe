/*
Copyright © 2024 NAME mazezen
*/
package cmd

import (
	"fmt"
	"github.com/mazezen/echoframe/cmd/auto/do"
	"github.com/mazezen/echoframe/cmd/auto/dt"
	"github.com/mazezen/itools"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// dtoCmd represents the dto command
var dtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runDaoCode()
	},
}

func init() {
	rootCmd.AddCommand(dtoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dtoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dtoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	handlerCmd.Flags().String("dto", "de", "generate dao code")
}

var df = "/internal/app/dao"

func runDaoCode() {
	dtoTem := rd()

	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strDto := itools.CompactStr(rootDir, df, strings.ToLower(model.Name), "_", "dto.go")
		if _, err := os.Stat(strDto); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}
	checkModel(ms)

	if err := generateDaoCode(ms, string(dtoTem)); err != nil {
		log.Fatal(err)
	}
}

func rd() []byte {
	dtoTem, err := os.ReadFile("./templates/dto.template")
	if err != nil {
		log.Fatal(err)
	}
	return dtoTem
}

func generateDaoCode(models []dt.ModelInfo, dtoTemplate string) error {
	dtoTmpl, err := template.New("dao").Parse(dtoTemplate)
	if err != nil {
		return err
	}

	sd := itools.CompactStr(rootDir, df)
	for _, model := range models {
		_, err = os.Stat(sd)
		if os.IsNotExist(err) {
			if err = os.Mkdir(sd, 0777); err != nil {
				log.Fatal(err)
			}
		}

		daoFileName := fmt.Sprintf("%s/%s_dao.go", sd, strings.ToLower(model.Name))
		_ = do.GenerateFile(dtoTmpl, daoFileName, model)
	}

	return nil
}
