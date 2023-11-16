package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type MxFile struct {
	XMLName xml.Name `xml:"mxfile"`
	Diagram Diagram
}

type Diagram struct {
	XMLName      xml.Name `xml:"diagram"`
	MxGraphModel MxGraphModel
}

type MxGraphModel struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    Root
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	MxCells []MxCell `xml:"mxCell"`
}

type MxCell struct {
	XMLName  xml.Name `xml:"mxCell"`
	Id       string   `xml:"id,attr"`
	Parent   string   `xml:"parent,attr"`
	Value    string   `xml:"value,attr"`
	Style    string   `xml:"style,attr"`
	Vertex   string   `xml:"vertex,attr"`
	Geometry MxGeomerty
}

type MxGeomerty struct {
	XMLName xml.Name `xml:"mxGeometry"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	As      string   `xml:"as,attr"`
}

// Return shapes found in XML File
func (mxFile *MxFile) GetShapes() []string {
	var known []string

	cells := mxFile.Diagram.MxGraphModel.Root.MxCells

	for i := 0; i < len(cells); i++ {
		style := cells[i].Style
		pattern := `;shape=(.*?);`
		re := regexp.MustCompile(pattern)
		submatches := re.FindStringSubmatch(style)
		if len(submatches) > 1 {
			shapeName := submatches[1]
			known = append(known, shapeName)
		}
	}
	return known
}

func Shapes(path string) ([]string, error) {
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Failed to open file: %s", path)
		os.Exit(1)
	}

	xmlFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	b, _ := io.ReadAll(xmlFile)

	var mxFile MxFile

	xml.Unmarshal(b, &mxFile)

	shapes := mxFile.GetShapes()

	return shapes, nil
}
