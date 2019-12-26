package main

import(
	"fmt"
	"os/exec"
	"strings"

)
func main(){
	// https://github.com/fogleman/primitive
	// the exec package takes two parameters the  command (here "primitive") and  the parameters
	cmd:=exec.Command("primitive", strings.Fields("-i input.png -o output.png -n 50 -m 6")...)
	b,err:=cmd.CombinedOutput()
	// b is the output and err is the error(if present) 
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))


}