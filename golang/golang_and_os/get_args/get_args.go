package get_args

import "os"

func GetArgs() []string {
	return os.Args[1:]
}
