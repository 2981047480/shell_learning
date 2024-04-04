package main

import (
	"fmt"

	"get_file_content"

	"github.com/tidwall/gjson"
)

func main() {
	// 解析json
	json := `{"name":"tidwall","age":18,"sex":"man"}`
	fmt.Println(gjson.Get(json, "name").String())
	fmt.Println(gjson.Get(json, "age").Int())
	fmt.Println(gjson.Get(json, "sex").String())
	fa := get_file_content.GetFactory()
	// str := "/Users/zephyrzhao/study/shell_learning/example_text/aaa"
	// str = "/Users/zephyrzhao/study/shell_learning/example_text/hosts"
	// fmt.Println(fa.Read_File("aaa"))
	a := "aaa"
	get_file_content.Panic_do(fa.Read_File, a)
}
