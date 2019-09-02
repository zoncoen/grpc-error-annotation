package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	errplugin "github.com/zoncoen/grpc-error-annotation/protoc-gen-go-grpc-error/plugin"
)

func main() {
	gen := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		gen.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "parsing input proto")
	}

	if len(gen.Request.FileToGenerate) == 0 {
		gen.Fail("no files to generate")
	}

	gen.CommandLineParameters(gen.Request.GetParameter())
	gen.WrapTypes()
	gen.SetPackageNames()
	gen.BuildTypeNameMap()
	gen.GeneratePlugin(errplugin.NewPlugin())

	for i := range gen.Response.File {
		gen.Response.File[i].Name = proto.String(strings.Replace(*gen.Response.File[i].Name, ".go", ".err.go", -1))
	}

	if err := goformat(gen.Response); err != nil {
		gen.Error(err)
	}
	data, err = proto.Marshal(gen.Response)
	if err != nil {
		gen.Error(err, "failed to marshal output proto")
	}
	if _, err = os.Stdout.Write(data); err != nil {
		gen.Error(err, "failed to write output proto")
	}
}

func goformat(resp *plugin.CodeGeneratorResponse) error {
	for i := 0; i < len(resp.File); i++ {
		formatted, err := format.Source([]byte(resp.File[i].GetContent()))
		if err != nil {
			return fmt.Errorf("failed to format file: %s", err)
		}
		fmts := string(formatted)
		resp.File[i].Content = &fmts
	}
	return nil
}
