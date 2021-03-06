package main

import "time"

var toolIndexs = []*Tool{
	&Tool{
		Name:      "studio",
		Alias:     "studio",
		Install:   "go get github.com/freezeChen/studio@" + Version,
		BuildTime: time.Date(2020, 4, 21, 23, 0, 0, 0, time.Local),
		Summary:   "studio工具集本体",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
	&Tool{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2020, 4, 21, 23, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
	&Tool{
		Name:      "genbts",
		Alias:     "studio-gen-bts",
		Install:   "go get github.com/freezeChen/studio/gen/studio-gen-bts@" + Version,
		BuildTime: time.Date(2020, 4, 21, 23, 0, 0, 0, time.Local),
		Summary:   "缓存回源逻辑代码生成器",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
	&Tool{
		Name:      "genrd",
		Alias:     "studio-gen-rd",
		Install:   "go get github.com/freezeChen/studio/gen/studio-gen-rd@" + Version,
		BuildTime: time.Date(2020, 4, 21, 23, 0, 0, 0, time.Local),
		Summary:   "redis缓存代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},

	&Tool{
		Name:      "protoc",
		Alias:     "studio-protoc",
		Install:   "go get github.com/freezeChen/studio/gen/studio-protoc@" + Version,
		BuildTime: time.Date(2020, 4, 21, 23, 0, 0, 0, time.Local),
		Summary:   "快速方便生成pb.go的protoc封装，windows、Linux请先安装protoc工具",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "studio",
	},
}
