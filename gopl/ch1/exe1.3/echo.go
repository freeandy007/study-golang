//做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异
package echo

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func echo2() {
	s, sep := "", ""
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}
func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
