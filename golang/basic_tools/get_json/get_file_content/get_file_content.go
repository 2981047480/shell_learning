package get_file_content

import (
	"fmt"
	"os"
)

type error_msg interface{}
type any interface{}
type FileContent interface {
	Write_File(filename any, content any) error
	Read_File(filename any) any
	PrintFileContent(filename any)
}

type Factory struct {
	f        *FileContent
	content  any
	filename any
}

func Panic_do(f func(any) (a any), a *any) any {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	res := f(a)
	return res
}

func GetFactory() *Factory {
	return new(Factory)
}

func (fac *Factory) Read_File(filename any) any {
	content, read_err := os.ReadFile(string(filename))
	if read_err != nil {
		// fmt.Println("Couldn't read file, error: ", read_err)
		// err_msg := "Couldn't read file, error: " + read_err.Error()
		// return "", err_msg
		panic("Couldn't read file, error: " + read_err.Error())
	}
	return any(content)
}

// func (fac *Factory) Read_File(filename any) (any, error_msg) {
// 	content, read_err := os.ReadFile(filename)
// 	if read_err != nil {
// 		// fmt.Println("Couldn't read file, error: ", read_err)
// 		// err_msg := "Couldn't read file, error: " + read_err.Error()
// 		// return "", err_msg
// 		panic("Couldn't read file, error: " + read_err.Error())
// 	}
// 	return any(content), nil
// }

func (f *Factory) PrintFileContent(filename any) {
	content, read_err := os.ReadFile(filename)
	if read_err != nil {
		fmt.Println("Couldn't read file, error: ", read_err)
		// return fa.content
	}
	fmt.Println(any(content))
}
