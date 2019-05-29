# studio
studio基于blibli开源项目[kratos](https://github.com/bilibili/kratos)根据自己的需求进行的修改与调整;  
包含一系列日常工作中所用到的各个项目工具集的管理与整合,以及项目模板代码的生成.

 
##安装工具
执行命令,快速安装studio工具
```
go get github.com/freezeChen/studio
```
##studio本体
```bash
NAME:
   studio - studio工具集

USAGE:
   studio [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
     new      :create template
     tool, t  studio tool
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

##studio new 
`studio new`用于快速生成项目模板代码  
模板代码基于go-micro框架搭建

```
studio new --name=demo

OPTIONS:
   --type value     web,srv (default: "web")
   --name value     app name
   --author value   author name (default: "freezeChen")
   --project value  single is false,use project as name
   --single         gen single project

```

```
├── main.go
│   └──

```

##studio tool
`studio tool`是基于protoc生成代码的工具集