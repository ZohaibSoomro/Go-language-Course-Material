package main

import(
	"fmt"
	"io"
	"os"
)

type Capper struct{
	writer io.Writer
}

func (c *Capper) Write(p []byte) (n int,err error){
	diff:=byte('a'-'A')
	out:=make([]byte,len(p))
	for i, c := range p {
		if c>='a' && c<='z'{	
			c-=diff
		}
		out[i]=c
	}
	return c.writer.Write(out)
}

func main(){
	cpr:=&Capper{os.Stdout}
	fmt.Fprintln(cpr,"Hello bro")
	os.Stdout.Write([]byte("Hi bro"))

}