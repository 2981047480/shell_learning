package pid

import "os"

func GetPid() (pid int) {
	return os.Getpid()
}
