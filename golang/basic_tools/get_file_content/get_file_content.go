package main

import (
	"fmt"
	"os"
)

type FileContent interface {
	GetFileContent(filename string) string
	PrintFileContent(filename string)
}

type Factory struct {
	f        *FileContent
	content  string
	filename string
}

func GetFactory() *Factory {
	return new(Factory)
}

func (fac *Factory) GetFileContent(filename string) string {
	content, read_err := os.ReadFile(filename)
	if read_err != nil {
		fmt.Println("Couldn't read file, error: ", read_err)
		// return fa.content
	}
	return string(content)
}

func (f *Factory) PrintFileContent(filename string) {
	content, read_err := os.ReadFile(filename)
	if read_err != nil {
		fmt.Println("Couldn't read file, error: ", read_err)
		// return fa.content
	}
	fmt.Println(string(content))
}

func main() {
	fa := GetFactory()
	fa.filename = "/Users/zephyrzhao/study/shell_learning/example_text/aaa"
	fmt.Println(fa.GetFileContent(fa.filename))
	fa.PrintFileContent(fa.filename)
}
