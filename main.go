package main

import(
	// "fmt"
	 "os"
	// "strings"
	"transform"
	"io"

)


func main(){
	inFile,err:=os.Open("input.png")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	out,err:=primitive.Transform(inFile,50)
	// out is the output and err is the error(if present) 
	if err != nil {
		panic(err)
	}
	os.Remove("out.png")
	outFile,err:=os.Create("out.png")
	if err != nil {
		panic(err)
	}
	io.Copy(outFile,out)


}