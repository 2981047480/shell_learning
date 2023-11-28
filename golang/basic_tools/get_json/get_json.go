package main

import (
	"fmt"

	"get_file_content"

	"github.com/tidwall/gjson"
)

// type

func main() {
	// 解析json
	fa := get_file_content.GetFactory()
	file := "/Users/zhaozephyr/shell_learning/shell_learning/example_text/aaa"
	content := fa.Read_File(file)
	json := `{"info":[{"name":"dj", "age":18},{"phone":"123456789","email":"dj@example.com"}]}`
	fmt.Println((content))
	fmt.Println(gjson.Get(content, "friends.#").String())                    // #用来显示元素的长度
	fmt.Println(gjson.Get(content, "friends.0").String())                    // 用来打印friends里面的第一个元素
	fmt.Println(gjson.Get(content, "frie*.#").String())                      // 先匹配，在访问长度
	fmt.Println(gjson.Get(content, "friends.#.first").String())              // 取到friends里面的所有键为first的值，并组成一个数组
	fmt.Println(gjson.Get(content, "friends.1.last").String())               // 取到friends里面的第二个元素的last的值
	fmt.Println(gjson.Get(content, "friends.1.age").String())                // 取age这个元素
	fmt.Println(gjson.Get(content, `friends.#(last="Murphy").age`).String()) // 匹配last，只匹配第一个
	fmt.Println(gjson.Get(content, `friends.#(age>45).first`).String())      // 匹配年龄大于45岁的，打印first name
	fmt.Println(gjson.Get(content, `friends.#(first%"D*"))`).String())       // 匹配friends里面first以D开头的元素，这里没有.什么，所以打印的是完整的值
	fmt.Println(gjson.Get(content, `friends.#(first!%"D*"))`).String())      // 匹配friends里面first不以D开头的第一个元素，这里没有.什么，所以打印的是完整的值
	fmt.Println(gjson.Get(content, `friends.#(nets.#(=="ig"))#`).String())   // 匹配friends里面first中nets含有ig的元素，这里没有.什么，所以打印的是完整的值
	fmt.Println(gjson.Get(content, `friends.#.nets|@flatten`).String())      // 不加flatten的效果：[["ig", "fb", "tw"],["fb", "tw"],["ig", "tw"]]  加了flatten的效果：["ig", "fb", "tw","fb", "tw","ig", "tw"]
	fmt.Println(gjson.Get(content, `friends|@reverse|0`).String())           // 先反转，再取第一个元素
	fmt.Println(gjson.Get(json, `info|@join`).String())                      // 把多个字典合并成一个
	// fmt.Println(gjson.Get(content, `@this|@pretty`).String())
	// fmt.Println(gjson.Get(content, `friends|@this`).String())
	// fmt.Println(gjson.Get(content, `friends|@valid`).String())
	/*
		有点像管道符:
			@reverse：翻转一个数组；
			@ugly：移除 JSON 中的所有空白符；
			@pretty：使 JSON 更易用阅读；
			@this：返回当前的元素，可以用来返回根元素；
			@valid：校验 JSON 的合法性；
			@flatten：数组平坦化，即将["a", ["b", "c"]]转为["a","b","c"]；
			@join：将多个对象合并到一个对象中。
	*/
}
