/*
   @Time : 2019-05-09 10:25
   @Author : frozenchen
   @File : new
   @Software: studio
*/
package new

import (
	"fmt"
	tmpl "github.com/freezeChen/studio/template"
	"github.com/freezeChen/studio/util"
	"github.com/micro/cli"
	"github.com/xlab/treeprint"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type config struct {
	// foo
	Alias string
	// micro new example -type
	Command string
	// go.micro
	Namespace string
	// api, srv, web, fnc
	Type string
	// go.micro.srv.foo
	FQDN string
	// github.com/micro/foo
	Dir string
	// $GOPATH/src/github.com/micro/foo
	GoDir string
	// $GOPATH
	GoPath string
	// Files
	Files []file
	// Comments
	Comments []string
	// Plugins registry=etcd:broker=nats
	Plugins []string
	Appname string
	Time    string
	Author  string
}

type file struct {
	Path string
	Tmpl string
}

func Command() []cli.Command {
	return []cli.Command{
		{
			Name:  "new",
			Usage: "create template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type",
					Usage: "web,srv",
					Value: "web",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "app name",
				},
				cli.StringFlag{
					Name:  "author",
					Usage: "author name",
					Value: "frozenChen",
				},
			},
			Action: func(context *cli.Context) {
				run(context)
			},
		},
	}
}

func run(ctx *cli.Context) {
	tye := ctx.String("type")
	appname := ctx.String("name")
	author := ctx.String("Author")
	var c config
	switch tye {
	case "web":

		c = config{
			Alias:   appname,
			FQDN:    "test",
			GoDir:   util.GetCurrentDirectory() + "/" + appname,
			Appname: appname,
			Author:  author,
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Files: []file{
				{"cmd/main.go", tmpl.Main_web},
				{"dao/dao.go", tmpl.Dao_template},
				{"go.mod", tmpl.Go_mod},
				{"model/model.go", tmpl.Model_template},
				{"server/http/http.go", tmpl.HttpServer_template},
				{"service/service.go", tmpl.Service_template},
				{"Makefile", tmpl.Makefile_template},
				{"conf/conf.go", tmpl.ConfDemo_template},
			},
			Comments: []string{
				"\ndownload protobuf for micro:\n",
				"brew install protobuf",
				"go get -u github.com/golang/protobuf/{proto,protoc-gen-go}",
				"go get -u github.com/micro/protoc-gen-micro",
				"\ncompile the proto file example.proto:\n",
				//"cd " + goDir,
				"protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto\n",
			},
		}
	}

	err := create(c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func create(c config) error {
	// check if dir exists
	if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.GoDir)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	fmt.Printf("Creating service %s in %s\n\n", c.FQDN, c.GoDir)

	t := treeprint.New()

	nodes := map[string]treeprint.Tree{}
	nodes[c.GoDir] = t

	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.GoDir, file.Path)
		dir := filepath.Dir(f)
		fmt.Println(dir)
		b, ok := nodes[dir]
		if !ok {
			d, _ := filepath.Rel(c.GoDir, dir)
			b = t.AddBranch(d)
			nodes[dir] = b
		}

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("mkdir err" + err.Error())
				return err
			}
		}

		p := filepath.Base(f)

		b.AddNode(p)
		if err := write(c, f, file.Tmpl); err != nil {
			fmt.Println("writeerr" + err.Error())
			return err
		}
	}

	// print tree
	fmt.Println(t.String())

	for _, comment := range c.Comments {
		fmt.Println(comment)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	return nil
}

func write(c config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": strings.Title,
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}
