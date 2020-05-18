package main

import (
	"errors"
	"flag"

	"github.com/zippunov/protorpc/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags        flag.FlagSet
		plugins      = flags.String("plugins", "", "deprecated option")
		importPrefix = flags.String("import_prefix", "", "deprecated option")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-gorpc: plugins are not supported;")
		}
		if *importPrefix != "" {
			return errors.New("protoc-gen-gorpc: import_prefix is not supported")
		}
		for _, f := range gen.Files {
			if f.Generate {
				plugin.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
