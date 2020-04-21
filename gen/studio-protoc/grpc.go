package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

const (
	_getGRPCGen = "go get github.com/gogo/protobuf/protoc-gen-gogofaster"

	_getMICROGen = "go get -u github.com/micro/protoc-gen-micro/v2"

	_grpcProtoc = "protoc --proto_path=. --proto_path=%s --micro_out=. --gogofaster_out=plugins=grpc:."


	//_grpcProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --gogofast_out=plugins=grpc:."
)

func installGRPCGen() error {
	if _, err := exec.LookPath("protoc-gen-gogofast"); err != nil {
		if err := goget(_getGRPCGen); err != nil {
			return err
		}
	}
	if _, err := exec.LookPath("protoc-gen-micro"); err != nil {
		if err := goget(_getMICROGen); err != nil {
			return err
		}
	}
	return nil
}

func genGRPC(ctx *cli.Context) error {
	return generate(ctx, fmt.Sprintf(_grpcProtoc,os.Getenv("GOPATH")+"/src"))
}
