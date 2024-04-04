package args_flag

import (
	"flag"
	"fmt"
)

func Args_flag() {
	str_arg := flag.String("str_arg", "init", "A string")
	int_arg := flag.Int("int_arg", 10, "An int")
	fmt.Println(str_arg, int_arg)
}
