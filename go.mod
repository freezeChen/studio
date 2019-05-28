module github.com/freezeChen/studio

go 1.12

require (
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/fatih/color v1.7.0
	github.com/freezeChen/studio-library v0.0.2
	github.com/hashicorp/go-msgpack v0.5.4 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	github.com/urfave/cli v1.20.0
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca
)

replace github.com/freezeChen/studio-library => ../studio-library
