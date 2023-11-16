package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tatomoaki/tfconfigbuilder/pkg/parser"
	"github.com/tatomoaki/tfconfigbuilder/pkg/resources"
	"github.com/tatomoaki/tfconfigbuilder/pkg/writer"
)

var rootCmd = &cobra.Command{
	Use:   "oss",
	Short: "Oss generate Terraform configuration files from diagrams.net drawings",
	Long:  ``,
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate configuration file from xml",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error %v", err)
		}

		shapes, _ := parser.Shapes(path)
		known := []resources.Resource{}
		for _, shape := range shapes {
			if resource, ok := resources.Resources[shape]; ok {
				known = append(known, resource)
			}
		}
		writer := writer.NewWriter()
		writer.Write(known)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
}

func init() {
	var resourceFile string
	genCmd.Flags().StringVarP(&resourceFile, "file", "f", "", "Path to xml file")
	genCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(genCmd)
}
