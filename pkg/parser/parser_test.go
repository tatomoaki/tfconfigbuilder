package parser_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/tatomoaki/tfconfigbuilder/pkg/parser"
)

func TestGetShapes(t *testing.T) {

	tests := []struct {
		pattern string
		content string
	}{
		{
			"temp.xml",
			`<?xml version="1.0" encoding="UTF-8"?>
			<mxfile>
			  <diagram>
				<mxGraphModel>
				  <root>
					<mxCell value="Circle" style="html=1;shape=circle;" vertex="1" parent="1">
					  <mxGeometry x="270" y="180" width="76.5" height="93" as="geometry" />
					</mxCell>
					<mxCell value="Rectangle" style="html=1;shape=rectangle;" vertex="1" parent="1">
					  <mxGeometry x="270" y="180" width="76.5" height="93" as="geometry" />
					</mxCell>
					<mxCell value="Square" style="html=1;shape=square;" vertex="1" parent="1">
					  <mxGeometry x="270" y="180" width="76.5" height="93" as="geometry" />
					</mxCell>
				  </root>
				</mxGraphModel>
			  </diagram>
			</mxfile>`,
		},
	}

	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			f, err := os.CreateTemp("", test.pattern)
			if err != nil {
				t.Errorf("TempFile(..., %q) error: %v", test.pattern, err)
			}
			f.Write([]byte(test.content))
			expected := []string{"circle", "rectangle", "square"}
			defer os.Remove(f.Name())

			got, err := parser.Shapes(f.Name())

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Expected file path %s, got %s ", expected, got)
			}

		})
	}
}
