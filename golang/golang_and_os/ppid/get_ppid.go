package ppid

import "os"

func GetPpid() (pid int) {
	return os.Getppid()
}
