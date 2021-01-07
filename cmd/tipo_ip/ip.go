package main
import (
	"net"
	"os"
	"fmt"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	//IPADDR
	addr := net.ParseIP(name) //recebe string
	if addr == nil {
		fmt.Println("Indereço invalido")
	} else {
		fmt.Println("O endereço é ", addr.String())
	}
	os.Exit(0)
}