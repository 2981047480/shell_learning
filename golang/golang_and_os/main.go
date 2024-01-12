package main

import (
	"fmt"

	"example.com/m/v2/golang_and_os/get_args"
	"example.com/m/v2/golang_and_os/pid"
	"example.com/m/v2/golang_and_os/ppid"
)

func main() {
	fmt.Println("pid", pid.GetPid(), "ppid", ppid.GetPpid(), "args", get_args.GetArgs())
}
