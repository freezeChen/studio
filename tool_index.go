package main

import "time"

var toolIndexs = []*Tool{
	&Tool{
		Name:      "studio",
		Alias:     "studio",
		BuildTime: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/freezeChen/studio",
		Summary:   "studio工具集本体",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
	&Tool{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2019, 5, 5, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
	&Tool{
		Name:      "genbts",
		Alias:     "kratos-gen-bts",
		BuildTime: time.Date(2019, 5, 5, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/bilibili/kratos/tool/kratos-gen-bts",
		Summary:   "缓存回源逻辑代码生成器",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
	&Tool{
		Name:      "genrd",
		Alias:     "studio-gen-rd",
		BuildTime: time.Date(2019, 5, 5, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/bilibili/kratos/tool/kratos-gen-mc",
		Summary:   "redis缓存代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
	&Tool{
		Name:      "protoc",
		Alias:     "studio-protoc",
		BuildTime: time.Date(2019, 5, 4, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/freezeChen/studio/gen/studio-protoc",
		Summary:   "快速方便生成pb.go的protoc封装，windows、Linux请先安装protoc工具",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "freeze",
	},
}
