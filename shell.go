package go_utils

import (
	"fmt"
	"os/exec"
)

//CallCommand
func CallCommand(cmd string) error {
	res := exec.Command(cmd)
	return res.Run()
	// p := procs.NewProcess(cmd)
	// p.OutputHandler = func(line string) string {
	// return line
	// }
	// p.ErrHandler = func(line string) string {
	// return line
	// }
	// return p.Run()
}

//go fmt
func GoFmt(file string) {
	CallCommand(fmt.Sprintf("go fmt %q", file))
}
