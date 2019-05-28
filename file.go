package main

import (
	"fmt"
	"os"
	"text/template"
)

type file struct {
	Path string
	Temp string
}

//func create(c config) error {
//	if _, err := os.Stat(getPath()); !os.IsNotExist(err) {
//		//return fmt.Errorf("%s already exists", c.getPath())
//	}
//
//	for _, file := range Files {
//
//		f := filepath.Join(getFilePath(), file.Path)
//		dir := filepath.Dir(f)
//
//		if !Single {
//			switch file.Path {
//			case "proto/hello.proto", "Makefile", "go.mod":
//				f = filepath.Join(util.GetCurrentDirectory()+"/"+Project, file.Path)
//				dir = filepath.Dir(f)
//
//			}
//		}
//
//		if _, err := os.Stat(dir); os.IsNotExist(err) {
//			if err := os.MkdirAll(dir, 0755); err != nil {
//				fmt.Println("mkdir err" + err.Error())
//				return err
//			}
//		}
//
//		if err := write(c, f, file.Temp); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func write(c config, file, temp string) error {
	fmt.Println(file)
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return nil
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	t, err := template.New("f").Parse(temp)
	if err != nil {
		return err
	}
	//if !c.Single {
	//	c.Appname = c.Project + "/" + c.Appname
	//}
	return t.Execute(f, c)
}
