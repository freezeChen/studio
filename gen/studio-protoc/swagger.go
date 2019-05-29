package main

import (
	"os/exec"

	"github.com/urfave/cli"
)

const (
	_getSwaggerGen = "go get github.com/bilibili/kratos/tool/protobuf/protoc-gen-bswagger"
	_swaggerProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bswagger_out=explicit_http=true:."
)

func installSwaggerGen() error {
	if _, err := exec.LookPath("protoc-gen-bswagger"); err != nil {
		if err := goget(_getSwaggerGen); err != nil {
			return err
		}
	}
	return nil
}

func genSwagger(ctx *cli.Context) error {
	return generate(ctx, _swaggerProtoc)
}
