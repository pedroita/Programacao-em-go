package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	if len(os.Args)!=2{
		fmt.Fprint(os.Stderr,"uso: %s hostname\n",os.Args[0])
		os.Exit(1)
	}
	addr,err := net.ResolveIPAddr("ip",os.Args[1])
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
