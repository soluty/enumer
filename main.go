package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var enumTemple = `package {{.Package}}

//go:generate stringer -type={{.Enum}}
type {{.Enum}} int

func (e {{.Enum}}) IsValid() bool {
	return e > _{{.Enum}}_Unknown && e < _{{.Enum}}_UnknownEnd
}

const (
	_{{.Enum}}_Unknown {{.Enum}} = iota


	_{{.Enum}}_UnknownEnd
)`

type Config struct {
	Output  string
	Package string
	Enum    string
}

var C = &Config{}

func main() {
	flag.StringVar(&C.Output, "o", "", "导出的文件名")
	flag.StringVar(&C.Package, "p", "enum", "导出文件所在包名")
	flag.StringVar(&C.Enum, "e", "", "枚举名")

	flag.Parse()

	if C.Enum == "" {
		log.Fatal("请输入要生成的枚举名称")
	}
	t, err := template.New("enum").Parse(enumTemple)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	enumName := C.Enum
	err = t.Execute(buf, struct {
		Enum    string
		Package string
	}{enumName, C.Package})
	if err != nil {
		log.Fatal(err)
	}
	if C.Output == "" {
		fmt.Println(buf.String())
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Getwd:", err)
	}
	outputFile := C.Output
	if !filepath.IsAbs(C.Output) {
		outputFile = filepath.Join(wd, outputFile)
	}
	if !strings.HasSuffix(outputFile, ".go") {
		outputFile = outputFile + ".go"
	}
	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("openfile:", err)
	}
	_, err = fmt.Fprintln(file, buf.String())
	if err == nil {
		fmt.Println("enum", C.Enum, "has write in file", outputFile)
	}
}
