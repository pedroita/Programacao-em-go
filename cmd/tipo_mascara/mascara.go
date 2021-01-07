package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main()  {
	if len(os.Args) !=4{
		fmt.Fprint(os.Stderr,"Uso : %s ip qtd-bists-mascara",os.Args[0])
		os.Exit(1)
	}
	addr := net.ParseIP(os.Args[1])
	if addr == nil{
		fmt.Println("Endereço invalido")
		os.Exit(1)
	}
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])
	//IPMask
	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)
	fmt.Println("O endereço é: ", addr.String(),
		"\nMáscara em hex: ",mask.String(),
		"\nRede é: ",network.String())
	os.Exit(0)



}
