package main

import (
	"bufio"
	"fmt"

	// "go/scanner"
	"os"
)

type any interface{}
type Scan interface {
	Scan() string
}

type Scanner struct {
	Text string
}

func (s *Scanner) Scan() []string {
	scan := bufio.NewScanner(os.Stdin)
	// dict := make(map[string]any)
	list := make([]string, 0)
	for scan.Scan() {
		if scan.Text() == "EOF" {
			break
		}
		list = append(list, scan.Text())
	}
	return list
}

func Get_Scanner() *Scanner {
	return new(Scanner)
}

func main() {
	// scan := bufio.NewScanner(os.Stdin)
	// // dict := make(map[string]any)
	// list := make([]string, 0)
	// for scan.Scan() {
	// 	if scan.Text() == "EOF" {
	// 		break
	// 	}
	// 	list = append(list, scan.Text())
	// }
	my_scanner := Get_Scanner()
	list := my_scanner.Scan()
	for _, v := range list {
		fmt.Println(v)
	} 
}
