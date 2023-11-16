package writer

import (
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tatomoaki/tfconfigbuilder/pkg/resources"
	"github.com/zclconf/go-cty/cty"
)

type Writer struct {
	File *hclwrite.File
}

// NewWriter returns a new Writer
func NewWriter() *Writer {
	file := hclwrite.NewEmptyFile()

	wr := &Writer{
		File: file,
	}
	return wr
}

// Write writes the resources to a file
func (w *Writer) Write(resources []resources.Resource) {
	body := w.File.Body()
	tfBlock := body.AppendNewBlock("terraform", nil)
	tfBlockBody := tfBlock.Body()

	requiredProvidersBlock := tfBlockBody.AppendNewBlock("required_providers", nil)
	requiredProvidersBlockBody := requiredProvidersBlock.Body()
	requiredProvidersBlockBody.SetAttributeValue("aws", cty.ObjectVal(map[string]cty.Value{
		"source":  cty.StringVal("hashicorp/aws"),
		"version": cty.StringVal("~> 5.0"),
	}))
	body.AppendNewline()
	provider := body.AppendNewBlock("provider", []string{"aws"})
	providerBody := provider.Body()
	providerBody.SetAttributeValue("region", cty.StringVal("af-south-1"))

	for _, resource := range resources {
		body.AppendNewline()
		resourceBlock := body.AppendNewBlock("resource", []string{resource.Name, "this"})
		resourceBody := resourceBlock.Body()
		for k, v := range resource.Attributes {
			resourceBody.SetAttributeValue(k, cty.StringVal(v))
		}
	}

	tFile, _ := os.Create("main.tf")
	tFile.Write(w.File.Bytes())

}
