package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type file struct {
	Path string
	Temp string
}

func create(c config)    error{
	fmt.Println(c.getPath())
	if _, err := os.Stat(c.getPath());!os.IsNotExist(err){
		return fmt.Errorf("%s already exists",c.getPath())
	}

	for _, file := range c.Files {
		f := filepath.Join(c.getPath(), file.Path)
		dir := filepath.Dir(f)
		if _, err := os.Stat(dir);os.IsNotExist(err){
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("mkdir err"+err.Error())
				return err
			}
		}


		if err := write(c, f, file.Temp); err != nil {
			return err
		}
	}

	return nil
}

func write(c config,file, temp string)error  {
	fmt.Println(file)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	t, err := template.New("f").Parse(temp)
	if err != nil {
		return err
	}
	return t.Execute(f,c)
}